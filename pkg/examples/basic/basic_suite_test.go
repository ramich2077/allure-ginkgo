package basic_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"allure-ginkgo/pkg/allure"
)

func TestBasic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basic Suite")
}

var _ = ReportAfterSuite("This description does nothing", func(report Report) {
	_ = allure.FromGinkgoReport(report)
})
