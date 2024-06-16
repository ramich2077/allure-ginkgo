package labels_test

import (
	. "github.com/onsi/ginkgo/v2"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

var _ = Describe("Severity", func() {
	It("Blocker", allure.Severity(allure.Blocker), func() {
		examples.Delay()
	})

	It("Critical", allure.Severity(allure.Critical), func() {
		examples.Delay()
	})

	It("Normal", allure.Severity(allure.Normal), func() {
		examples.Delay()
	})

	It("Minor", allure.Severity(allure.Minor), func() {
		examples.Delay()
	})

	It("Trivial", allure.Severity(allure.Trivial), func() {
		examples.Delay()
	})
})

var _ = Describe("BDD Labels", func() {
	It("Dummy test",
		allure.Story("New client registers"),
		allure.Feature("SSO log-on"),
		allure.Epic("New client"),
		func() {
			examples.Delay()
		})
})

var _ = Describe("Tags", func() {
	It("Tagged test", allure.Tag("Tag One"), allure.Tag("Tag Two"), func() {
		examples.Delay()
	})
})

var _ = Describe("Custom suite", func() {
	It("Test with custom suite", allure.Suite("Custom suite"), func() {
		examples.Delay()
	})
})

var _ = Describe("Owner", func() {
	It("Test with marked owner", allure.Owner("Ginkgo"), func() {
		examples.Delay()
	})
})

var _ = Describe("Id", func() {
	It("Test with marked owner", allure.Id("123"), func() {
		examples.Delay()
	})
})

var _ = Describe("Custom label", func() {
	It("Test with custom label", allure.CustomLabel("test-image", "some-image:latest"), func() {
		examples.Delay()
	})
})
