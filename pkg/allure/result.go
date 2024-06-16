package allure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/rs/zerolog/log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"github.com/onsi/ginkgo/v2/types"
)

const descriptionReportEntryName = "DESCRIPTION"

// result is the top level report object for a test.
type result struct {
	UUID          string         `json:"uuid,omitempty"`
	TestCaseID    string         `json:"testCaseId,omitempty"`
	HistoryID     string         `json:"historyId,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Status        string         `json:"status,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage,omitempty"`
	Steps         []stepObject   `json:"steps,omitempty"`
	Attachments   []attachment   `json:"attachments,omitempty"`
	Parameters    []parameter    `json:"parameters,omitempty"`
	Start         int64          `json:"start,omitempty"`
	Stop          int64          `json:"stop,omitempty"`
	Children      []string       `json:"children,omitempty"`
	FullName      string         `json:"fullName,omitempty"`
	Labels        []label        `json:"labels,omitempty"`
	Links         []Link         `json:"links,omitempty"`
	Suite         string         `json:"-"`
	ParentSuite   string         `json:"-"`
}

func (r *result) addSeverity(severity SeverityLevel) *result {
	r.addLabel(labelSeverity, string(severity))

	return r
}

func (r *result) addSuite(suite string) *result {
	r.Suite = suite
	r.addLabel(labelSuite, suite)

	return r
}

func (r *result) addParentSuite(parentSuite string) *result {
	r.ParentSuite = parentSuite
	r.addLabel(labelParentSuite, parentSuite)

	return r
}

func (r *result) addParameters(parameters map[string]string) *result {
	for key, value := range parameters {
		r.Parameters = append(r.Parameters, parameter{Name: key, Value: value})
	}

	return r
}

func (r *result) addAttachment(attachment *attachment) *result {
	if attachment == nil {
		log.Err(fmt.Errorf("nil attachment pointer"))
		return r
	}

	r.Attachments = append(r.Attachments, *attachment)

	return r
}

func (r *result) addFullName(FullName string) *result {
	r.FullName = FullName

	return r
}

func (r *result) addLabel(name string, value string) *result {
	r.Labels = append(r.Labels, label{
		Name:  name,
		Value: value,
	})

	return r
}

func (r *result) addLinks(links ...Link) *result {
	r.Links = append(r.Links, links...)

	return r
}

func (r *result) setStatusDetails(details statusDetails) *result {
	r.StatusDetails = &details

	return r
}

func (r *result) createFromSpecReport(specReport ginkgo.SpecReport) *result {
	r.Start = getTimestampMsFromTime(specReport.StartTime)
	r.Stop = getTimestampMsFromTime(specReport.EndTime)

	if r.Stop < r.Start { //Workaround for incorrect skipped tests execution time
		r.Stop = r.Start
	}

	r.Name = specReport.LeafNodeText
	r.Description = buildDescription(specReport)

	r.setDefaultLabels(specReport)

	if len(specReport.ContainerHierarchyTexts) > 0 {
		r.addSuite(specReport.ContainerHierarchyTexts[len(specReport.ContainerHierarchyTexts)-1])
	} else {
		r.addSuite(r.Name)
	}

	r.setAllureLabels(specReport)

	attachmentEntries := filterForAttachments(specReport.ReportEntries)
	var toSkip map[int]struct{}
	r.Steps, toSkip = createSteps(specReport.SpecEvents, attachmentEntries)

	for i, entry := range attachmentEntries {
		if _, ok := toSkip[i]; !ok {

			var att attachment
			err := json.Unmarshal([]byte(entry.Value.GetRawValue().(string)), &att)

			if err != nil {
				log.Err(fmt.Errorf("error processing attachment for entry %s on line %d", entry.Location.FileName, entry.Location.LineNumber))
				continue
			} else if reflect.DeepEqual(att, attachment{}) {
				log.Err(fmt.Errorf("nil pointer attachment for entry %s on line %d", entry.Location.FileName, entry.Location.LineNumber))
				continue
			}

			r.addAttachment(&att)
		}
	}

	currentHash := uuid.NewSHA1(
		uuid.Nil, []byte(strings.Join([]string{r.Name, r.Suite, r.ParentSuite}, ""))).String()
	r.TestCaseID = currentHash
	r.HistoryID = currentHash

	r.Stage = "finished"
	r.Status = getTestStatus(specReport)

	if r.Status == failed || r.Status == broken {
		details := statusDetails{
			Message: specReport.Failure.Message,
			Trace:   specReport.Failure.Location.FullStackTrace,
		}
		r.setStatusDetails(details)
	}

	return r
}

func createSteps(events types.SpecEvents, entries types.ReportEntries) (steps []stepObject, indicesToSkip map[int]struct{}) {
	currentEndIndex := -1
	indicesToSkip = make(map[int]struct{})
	steps = []stepObject{}

	for startEventIndex, startEvent := range events {
		if currentEndIndex >= startEventIndex {
			//Skipping all nested steps from previous iterations
			continue
		}

		if startEvent.SpecEventType == types.SpecEventByStart {
			step := newStep()
			step.addName(startEvent.Message)
			step.Status = passed
			step.Stage = "finished"
			endEvent, endIndex := findByEventEnd(events, startEvent)

			if endEvent != nil {
				step.Start = getTimestampMsFromTime(startEvent.TimelineLocation.Time)
				step.Stop = getTimestampMsFromTime(endEvent.TimelineLocation.Time)

				childrenSteps, toSkip := createSteps(events[startEventIndex+1:endIndex], entries)

				step.ChildrenSteps = childrenSteps

				for i, entry := range entries {
					if _, ok := toSkip[i]; !ok {
						if entry.TimelineLocation.Order > startEvent.TimelineLocation.Order &&
							entry.TimelineLocation.Order < endEvent.TimelineLocation.Order {
							var att attachment
							err := json.Unmarshal([]byte(entry.Value.GetRawValue().(string)), &att)
							if err != nil {
								log.Err(fmt.Errorf("error processing attachment for entry %s on line %d", entry.Location.FileName, entry.Location.LineNumber))
								toSkip[i] = struct{}{}
								continue
							} else if reflect.DeepEqual(att, attachment{}) {
								log.Err(fmt.Errorf("nil pointer attachment for entry %s on line %d", entry.Location.FileName, entry.Location.LineNumber))
								toSkip[i] = struct{}{}
								continue
							}
							step.addAttachment(&att)

							toSkip[i] = struct{}{}
						}
					}
				}

				for k, v := range toSkip {
					indicesToSkip[k] = v
				}
				currentEndIndex = endIndex
			}

			steps = append(steps, *step)
		}
	}
	return steps, indicesToSkip
}

func findByEventEnd(events types.SpecEvents, startEvent types.SpecEvent) (event *types.SpecEvent, index int) {
	for i, e := range events {
		if e.SpecEventType == types.SpecEventByEnd &&
			startEvent.CodeLocation.LineNumber == e.CodeLocation.LineNumber &&
			startEvent.TimelineLocation.Order < e.TimelineLocation.Order {
			return &e, i
		}
	}

	return nil, -1
}

func filterForAttachments(entries types.ReportEntries) types.ReportEntries {
	var res types.ReportEntries
	for _, entry := range entries {
		if entry.Name == attachmentReportEntryName {
			res = append(res, entry)
		}
	}

	return res
}

func buildDescription(specReport ginkgo.SpecReport) string {
	containerDescs := make([]string, 0)
	if len(specReport.ContainerHierarchyTexts) > 1 {
		//every container text excluding the top-level suite desc
		containerDescs = append(containerDescs, specReport.ContainerHierarchyTexts[1:]...)
	}

	var nodeDesc string
	for _, entry := range specReport.ReportEntries {
		if entry.Name == descriptionReportEntryName {
			nodeDesc = entry.Value.GetRawValue().(string)
		}
	}

	return strings.Join(append(containerDescs, nodeDesc), "\n")
}

func (r *result) setDefaultLabels(report ginkgo.SpecReport) *result {
	wsd := os.Getenv(wsPathEnvKey)

	programCounters := make([]uintptr, 10)
	callersCount := runtime.Callers(0, programCounters)
	var testFile string
	for i := 0; i < callersCount; i++ {
		_, testFile, _, _ = runtime.Caller(i)
		if strings.Contains(testFile, "_test.go") {
			break
		}
	}
	testPackage := strings.TrimSuffix(strings.Replace(strings.TrimPrefix(testFile, wsd+"/"), "/", ".", -1), ".go")

	if report.IsSerial {
		r.addLabel("thread", "0")
	} else {
		r.addLabel("thread", strconv.Itoa(report.ParallelProcess))
	}

	r.addLabel("package", testPackage)
	r.addLabel("testClass", testPackage)
	r.addLabel("testMethod", report.LeafNodeText)
	if len(wsd) == 0 {
		r.addFullName(fmt.Sprintf("%s:%s", report.FileName(), report.LeafNodeText))
	} else {
		r.addFullName(fmt.Sprintf("%s:%s", strings.TrimPrefix(report.FileName(), wsd+"/"), report.LeafNodeText))
	}
	if hostname, err := os.Hostname(); err == nil {
		r.addLabel("host", hostname)
	}

	r.addLabel("language", "golang")

	return r
}

func (r *result) setAllureLabels(specReport ginkgo.SpecReport) *result {
	for _, ginkgoLabel := range specReport.Labels() {
		if strings.Contains(ginkgoLabel, allureLabelPrefix) {
			aLabel := strings.TrimPrefix(strings.TrimPrefix(ginkgoLabel, allureLabelPrefix), ".")
			switch kv := strings.SplitN(aLabel, ":", 2); kv[0] {
			case labelSeverity:
				if len(kv) > 1 {
					r.addSeverity(SeverityLevel(kv[1]))
				} else {
					r.addSeverity(Normal)
				}
			case labelParametrized:
				if len(kv) > 1 {
					bs, err := base64.URLEncoding.DecodeString(kv[1])
					if err != nil {
						log.Err(err)
					}

					var paramsMap map[string]string
					err = json.Unmarshal(bs, &paramsMap)
					if err != nil {
						log.Err(err)
					}

					r.addParameters(paramsMap)
				} else {
					log.Err(fmt.Errorf(
						"emprty parameters in parametrized label, test: %s, suite: %s", r.Name, r.ParentSuite))
				}
			case labelLink:
				if len(kv) > 1 {
					bs, err := base64.URLEncoding.DecodeString(kv[1])
					if err != nil {
						log.Err(err)
					}

					var link Link
					err = json.Unmarshal(bs, &link)
					if err != nil {
						log.Err(err)
					}

					r.addLinks(link)
				} else {
					log.Err(fmt.Errorf(
						"emprty parameters in parametrized label, test: %s, suite: %s", r.Name, r.ParentSuite))
				}
			default:
				if len(kv) > 1 {
					r.addLabel(kv[0], kv[1])
				} else {
					r.addLabel(kv[0], "")
				}
			}
		}
	}

	return r
}

func (r *result) write() {
	content, err := json.Marshal(r)
	if err != nil {
		log.Err(fmt.Errorf("failed to marshall result into MimeTypeJSON: %w", err))
	}

	err = writeFile(fmt.Sprintf("%s-result.json", r.TestCaseID), content)
	if err != nil {
		log.Err(fmt.Errorf("failed to write content of result to json file: %w", err))
	}
}

func newResult() *result {
	return &result{
		UUID:  uuid.New().String(),
		Start: getTimestampMs(),
		Steps: []stepObject{},
	}
}
