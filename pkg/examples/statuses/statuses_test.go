package statuses_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"allure-ginkgo/pkg/examples"
)

var _ = Describe("Test statuses suite", func() {

	It("Test passing", func() {
		examples.Delay()
		Expect(true).To(BeTrue())
	})

	It("Test failing", func() {
		examples.Delay()
		Expect(true).To(BeFalse())
	})

	XIt("Test skipped", func() {
		examples.Delay()
		Expect(true).To(BeTrue()) //this won't be executed
	})

	It("Test broken", func() {
		examples.Delay()
		panic(fmt.Sprintf("Broken test"))
	})

})
