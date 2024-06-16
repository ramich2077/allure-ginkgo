package allure

type label struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type SeverityLevel string

const (
	Blocker  SeverityLevel = "blocker"
	Critical SeverityLevel = "critical"
	Normal   SeverityLevel = "normal"
	Minor    SeverityLevel = "minor"
	Trivial  SeverityLevel = "trivial"
)

const (
	allureLabelPrefix = "allure"

	labelSeverity = "severity"

	labelEpic = "epic"

	labelStory = "story"

	labelFeature = "feature"

	labelTag = "tag"

	labelOwner = "owner"

	labelId = "AS_ID"

	labelSuite = "suite"

	labelSubSuite = "subSuite"

	labelParentSuite = "parentSuite"

	labelParametrized = "parametrized"

	labelLink = "link"

	labelLead = "lead"
)
