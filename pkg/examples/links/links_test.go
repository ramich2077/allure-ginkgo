package links_test

import (
	. "github.com/onsi/ginkgo/v2"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

var _ = It("Tests with links",
	allure.AddLink(
		"Generik link",
		"https://onsi.github.io/ginkgo",
		allure.LinkGeneric,
	),
	allure.AddLink(
		"Issue",
		"https://github.com/onsi/ginkgo/issues/1422",
		allure.LinkIssue,
	),
	allure.AddLink(
		"Pull request",
		"https://github.com/onsi/ginkgo/pull/1203",
		allure.LinkPr,
	),
	allure.AddLink(
		"Test case",
		"https://github.com/onsi/ginkgo/issues/1422",
		allure.LinkTms,
	),
	func() {
		examples.Delay()
	})
