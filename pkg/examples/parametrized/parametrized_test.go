package parametrized_test

import (
	. "github.com/onsi/ginkgo/v2"

	"allure-ginkgo/pkg/allure"
	"allure-ginkgo/pkg/examples"
)

type DummyInnerObj struct {
	Num  float64
	Flag bool
	Vals []any
	Dict map[string]any
}

type DummyObj struct {
	Id   int
	Name string
	Data DummyInnerObj
}

var (
	b     = false
	s     = "dummy string"
	i     = 12345
	f     = 1.2345
	slice = []any{"a", 1}
	kv    = map[string]any{
		"a": 1,
		"b": 1.0,
		"c": "dummy",
	}
	dummy = DummyObj{
		Id:   i,
		Name: s,
		Data: DummyInnerObj{
			Num:  f,
			Flag: b,
			Vals: slice,
			Dict: kv,
		},
	}
)

var _ = Describe("Parametrized tests suite", func() {
	DescribeTable("This is a parametrized test", func(args ...any) {
		By("Parametrized assertion", func() {
			examples.Delay()
		})
	},
		allure.AddEntry(Entry("first static description", b, s, i, f, slice, kv, dummy)),
		allure.AddEntry(Entry(EntryDescription("first format string, args: %b %s %d %f %v %+v %+v"), b, s, i, f, slice, kv, dummy)),
		allure.AddEntry(Entry(nil, b, s, i, f, slice, kv, dummy)),
		allure.AddEntry(
			Entry(func(args ...any) string { return "This is a first description from function" },
				b, s, i, f, slice, kv, dummy),
		),
	)
})

var _ = DescribeTable("This is test with named parameters",
	func(args ...any) {

		By("Parametrized assertion", func() {
			examples.Delay()
		})
	},
	allure.AddEntry(
		Entry("second static description",
			allure.NewParam("a bool", b),
			allure.NewParam("a string", s),
			allure.NewParam("an int", i),
			allure.NewParam("a float", f),
			allure.NewParam("a slice", slice),
			allure.NewParam("a map", kv),
			allure.NewParam("an obj", dummy),
		),
	),
	allure.AddEntry(
		Entry(EntryDescription("second format string, args: %b %s %d %f %v %+v %+v"),
			allure.NewParam("a bool", b),
			allure.NewParam("a string", s),
			allure.NewParam("an int", i),
			allure.NewParam("a float", f),
			allure.NewParam("a slice", slice),
			allure.NewParam("a map", kv),
			allure.NewParam("an obj", dummy),
		),
	),
	allure.AddEntry(
		Entry(nil,
			allure.NewParam("a bool", b),
			allure.NewParam("a string", s),
			allure.NewParam("an int", i),
			allure.NewParam("a float", f),
			allure.NewParam("a slice", slice),
			allure.NewParam("a map", kv),
			allure.NewParam("an obj", dummy),
		),
	),
	allure.AddEntry(
		Entry(func(args ...any) string { return "This is a second description from function" },
			allure.NewParam("a bool", b),
			allure.NewParam("a string", s),
			allure.NewParam("an int", i),
			allure.NewParam("a float", f),
			allure.NewParam("a slice", slice),
			allure.NewParam("a map", kv),
			allure.NewParam("an obj", dummy),
		),
	),
)
