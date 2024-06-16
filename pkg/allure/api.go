package allure

import (
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"
)

func Description(description string) {
	ginkgo.AddReportEntry(descriptionReportEntryName, ginkgo.ReportEntryVisibilityNever, ginkgo.Offset(1), description)
}

func AddAttachment(name string, mimeType MimeType, content []byte) {
	a, _ := addAttachment(name, mimeType, content)
	//Here we are marshalling the attachment object itself to JSON, so it can be transferred between parallel processes
	ginkgo.AddReportEntry(attachmentReportEntryName, ginkgo.ReportEntryVisibilityNever, ginkgo.Offset(1), string(saveAsJSONAttachment(&a)))
}

func Severity(severity SeverityLevel) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelSeverity, severity))
}

func Tag(tag string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelTag, tag))
}

func Owner(owner string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelOwner, owner))
}

func Epic(epic string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelEpic, epic))
}

func Feature(feature string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelFeature, feature))
}

func CustomLabel(label, value string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, label, value))
}

func Id(id string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelId, id))
}

func Story(story string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelStory, story))
}

func Suite(suite string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelSuite, suite))
}

func SubSuite(subSuite string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelSubSuite, subSuite))
}

func Lead(lead string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelLead, lead))
}

func Parametrized(value string) ginkgo.Labels {
	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelParametrized, value))
}

func AddEntry(entry ginkgo.TableEntry) ginkgo.TableEntry {
	ginkgo.GinkgoHelper()
	return addEntry(entry)
}

func NewParam(name string, value any) Parameter {
	ginkgo.GinkgoHelper()
	return Parameter{
		Name:  name,
		Value: value,
	}
}

func FromGinkgoReport(report types.Report) error {
	return newTestContainer().createFromReport(report).write()
}
