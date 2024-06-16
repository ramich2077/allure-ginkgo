package labels_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"allure-ginkgo/pkg/allure"
)

func TestLabels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Labels Suite")
}

var _ = ReportAfterSuite("This description does nothing", func(report Report) {
	_ = allure.FromGinkgoReport(report)
})
