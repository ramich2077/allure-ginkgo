package allure

import (
	"encoding/json"
	"fmt"
	"github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"os"
	"strings"

	"github.com/google/uuid"

	"github.com/pkg/errors"
)

type attachment struct {
	uuid    string
	Name    string   `json:"name"`
	Source  string   `json:"source"`
	Type    MimeType `json:"type"`
	content []byte
}

type MimeType string

const (
	MimeTypeJSON MimeType = "application/json"
	MimeTypeCSV  MimeType = "text/csv"
	MimeTypeText MimeType = "text/plain"
	MimeTypePNG  MimeType = "image/png"
	MimeTypeHTML MimeType = "text/html"
)

const attachmentReportEntryName = "ATTACHMENT"

func saveAsJSONAttachment(msg any) []byte {
	res, e := json.MarshalIndent(msg, "", "  ")
	gomega.Expect(e).ShouldNot(gomega.HaveOccurred(), "while marshalling raw message")
	return res
}

func addAttachment(name string, mimeType MimeType, content []byte) (*attachment, error) {
	attachment := newAttachment(name, mimeType, content)
	err := attachment.writeAttachmentFile()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create an attachment file")
	}

	return attachment, nil
}

func (a *attachment) writeAttachmentFile() error {
	resultsPathEnv := os.Getenv(resultsPathEnvKey)
	ensureFolderCreated()
	if resultsPathEnv == "" {
		log.Warn().Msg(fmt.Sprintf("%s environment variable cannot be empty\n", resultsPathEnvKey))
	}
	if resultsPath == "" {
		resultsPath = fmt.Sprintf("%s/allure-results", resultsPathEnv)
	}

	a.Source = fmt.Sprintf("%s-attachment.%s", a.uuid, resolveExtension(a.Type))
	err := os.WriteFile(strings.Join([]string{resultsPath, a.Source}, "/"), a.content, 0600)
	if err != nil {
		return fmt.Errorf("failed to write in file: %w", err)
	}

	return nil
}

func newAttachment(name string, mimeType MimeType, content []byte) *attachment {
	result := &attachment{
		uuid:    uuid.New().String(),
		content: content,
		Name:    name,
		Type:    mimeType,
	}

	return result
}

func resolveExtension(mimeType MimeType) string {
	switch mimeType {
	case MimeTypeJSON:
		return "json"
	case MimeTypeCSV:
		return "csv"
	case MimeTypeHTML:
		return "html"
	case MimeTypePNG:
		return "png"
	case MimeTypeText:
		return "txt"
	default:
		return ""
	}
}
