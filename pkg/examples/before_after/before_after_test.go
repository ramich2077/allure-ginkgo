package before_after_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"

	"allure-ginkgo/pkg/examples"
)

var counter = 0

var _ = Describe("Dummy test suite for setup & teardown demo", func() {

	BeforeEach(func() {
		By(fmt.Sprintf("Before each setup, counter value: %d", counter), func() {
			examples.Delay()
		})
	})

	AfterEach(func() {
		By(fmt.Sprintf("After each teardown, counter value: %d", counter), func() {
			examples.Delay()
		})
	})

	It("First test", func() {
		examples.Delay()
		counter = counter + 1
	})

	It("Second test", func() {
		examples.Delay()
		counter = counter + 1
	})

	It("Third test", func() {
		examples.Delay()
		counter = counter + 1
	})

})
