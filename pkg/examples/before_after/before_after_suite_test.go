package before_after_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

func TestBeforeAfter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BeforeAfter Suite")
}

var _ = SynchronizedBeforeSuite(
	func() {
		By("Synchronized before suite setup - singe process", examples.Delay)
		DeferCleanup(func() {
			By("Single process defer cleanup", examples.Delay)
		})
	},
	func() {
		By(fmt.Sprintf("Synchronized before suite setup - process %d", GinkgoParallelProcess()), examples.Delay)
		DeferCleanup(func() {
			By(fmt.Sprintf("Multiple processes defer cleanup - process %d", GinkgoParallelProcess()), examples.Delay)
		})
	})

var _ = SynchronizedAfterSuite(
	func() {
		By(fmt.Sprintf("Synchronized after suite teardown - process %d", GinkgoParallelProcess()), examples.Delay)
	},
	func() {
		By("Synchronized after suite teardown - singe process")
	})

var _ = ReportAfterSuite("This description does nothing", func(report Report) {
	_ = allure.FromGinkgoReport(report)
})
