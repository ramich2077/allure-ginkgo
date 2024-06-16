package allure

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

type stepObject struct {
	Name          string         `json:"name,omitempty"`
	Status        string         `json:"status,omitempty"`
	Description   string         `json:"description,omitempty"`
	StatusDetails *statusDetails `json:"statusDetails,omitempty"`
	Stage         string         `json:"stage"`
	ChildrenSteps []stepObject   `json:"steps"`
	Attachments   []attachment   `json:"attachments"`
	Parameters    []parameter    `json:"parameters"`
	Start         int64          `json:"start"`
	Stop          int64          `json:"stop"`
}

func (sc *stepObject) addName(name string) {
	sc.Name = name
}

func (sc *stepObject) addAttachment(attachment *attachment) {
	if attachment == nil {
		log.Err(fmt.Errorf("nil attachment pointer"))
		return
	}
	sc.Attachments = append(sc.Attachments, *attachment)
}

func newStep() *stepObject {
	return &stepObject{
		Attachments:   make([]attachment, 0),
		ChildrenSteps: make([]stepObject, 0),
		Parameters:    make([]parameter, 0),
	}
}
