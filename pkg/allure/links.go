package allure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/rs/zerolog/log"
)

type Link struct {
	Type LinkType `json:"type"`
	Url  string   `json:"url"`
	Name string   `json:"name"`
}

type LinkType string

const (
	LinkGeneric LinkType = "link"
	LinkIssue   LinkType = "issue"
	LinkTms     LinkType = "tms"
	LinkPr      LinkType = "pr"
)

func AddLink(name string, url string, linkType LinkType) ginkgo.Labels {
	ginkgo.GinkgoHelper()

	link := Link{
		Type: linkType,
		Url:  url,
		Name: name,
	}

	linkJson, err := json.Marshal(link)
	if err != nil {
		log.Err(err)
	}

	linkEncoded := base64.URLEncoding.EncodeToString(linkJson)

	return ginkgo.Label(fmt.Sprintf("%s.%s:%s", allureLabelPrefix, labelLink, linkEncoded))
}
