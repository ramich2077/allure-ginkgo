package attachments

import (
	_ "embed"
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

//go:embed resources/allure-logo.png
var pngExample []byte

var _ = Describe("Attachments suite", func() {
	It("Attachments test", func() {
		structJson, _ := json.Marshal(struct {
			ID         int
			Name       string
			Number     float64
			Collection []string
		}{
			ID:         1,
			Name:       "JSON example",
			Number:     1.0,
			Collection: []string{"first", "second"},
		})

		allure.AddAttachment("example.json", allure.MimeTypeJSON, structJson)

		allure.AddAttachment("example.txt", allure.MimeTypeText, []byte("Some arbitrary text"))

		allure.AddAttachment("example.csv", allure.MimeTypeCSV, []byte("first,second,third\nnone,two,three"))

		allure.AddAttachment("example.html", allure.MimeTypeHTML, []byte("<h1>Example html attachment</h1>"))

		allure.AddAttachment("example.png", allure.MimeTypePNG, pngExample)

		examples.Delay()
	})

	It("Attachments in steps test", func() {
		examples.Delay()

		By("Step with attachments", func() {
			structJson, _ := json.Marshal(struct {
				ID         int
				Name       string
				Number     float64
				Collection []string
			}{
				ID:         1,
				Name:       "JSON example",
				Number:     1.0,
				Collection: []string{"first", "second"},
			})

			allure.AddAttachment("example.json", allure.MimeTypeJSON, structJson)

			allure.AddAttachment("example.txt", allure.MimeTypeText, []byte("Some arbitrary text"))

			allure.AddAttachment("example.csv", allure.MimeTypeCSV, []byte("first,second,third\nnone,two,three"))

			allure.AddAttachment("example.html", allure.MimeTypeHTML, []byte("<h1>Example html attachment</h1>"))

			allure.AddAttachment("example.png", allure.MimeTypePNG, pngExample)

			By("Substep with attachments", func() {

				allure.AddAttachment("second_example.json", allure.MimeTypeJSON, structJson)

				allure.AddAttachment("second_example.txt", allure.MimeTypeText, []byte("Some arbitrary text"))

				allure.AddAttachment("second_example.csv", allure.MimeTypeCSV, []byte("first,second,third\nnone,two,three"))

				allure.AddAttachment("second_example.html", allure.MimeTypeHTML, []byte("<h1>Example html attachment</h1>"))

				allure.AddAttachment("example.png", allure.MimeTypePNG, pngExample)
			})
		})

	})

})
