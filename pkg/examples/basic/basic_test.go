package basic_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

var _ = Describe("Simple suite", func() {

	It("Simple test", func() {
		allure.Description("This is a simple test")

		examples.Delay()

		By("Simple step", func() {
			Expect("simple assert").To(Equal("simple assert"))
		})

		By("Simple parent step", func() {
			Expect("parent step assert").To(Equal("parent step assert"))
			By("Simple child step", func() {
				Expect("child step assert").To(Equal("child step assert"))
			})
		})

		By("Simple step without wrapping function")
		Expect("test body assert").To(Equal("test body assert"))
	})
})
