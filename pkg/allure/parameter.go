package allure

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/rs/zerolog/log"
	"reflect"
	"unsafe"
)

type parameter struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Parameter struct {
	Name  string `json:"name,omitempty"`
	Value any    `json:"value,omitempty"`
}

func addEntry(entry ginkgo.TableEntry) ginkgo.TableEntry {
	ginkgo.GinkgoHelper()
	v := reflect.ValueOf(entry)
	vTmp := reflect.New(v.Type()).Elem()
	vTmp.Set(v)

	descVal := vTmp.FieldByName("description")
	descVal = reflect.NewAt(descVal.Type(), unsafe.Pointer(descVal.UnsafeAddr())).Elem()
	decorsVal := vTmp.FieldByName("decorations")
	decorsVal = reflect.NewAt(decorsVal.Type(), unsafe.Pointer(decorsVal.UnsafeAddr())).Elem()
	paramsVal := vTmp.FieldByName("parameters")
	paramsVal = reflect.NewAt(paramsVal.Type(), unsafe.Pointer(paramsVal.UnsafeAddr())).Elem()

	description := descVal.Interface()
	decorations := decorsVal.Interface().([]any)
	parameters := paramsVal.Interface().([]any)

	paramsMap := make(map[string]string, len(parameters))
	for i, param := range parameters {
		if allureParam, ok := param.(Parameter); ok {
			paramsMap[allureParam.Name] = fmt.Sprintf("%+v", allureParam.Value)
			parameters[i] = allureParam.Value
		} else {
			paramsMap[fmt.Sprintf("arg%d", i)] = fmt.Sprintf("%+v", param)
		}
	}

	paramsJson, err := json.Marshal(paramsMap)
	if err != nil {
		log.Err(err)
	}

	paramsEncoded := base64.URLEncoding.EncodeToString(paramsJson)
	parametrizedLabel := Parametrized(paramsEncoded)
	decorations = append(decorations, parametrizedLabel)

	args := append(decorations, parameters...)

	return ginkgo.Entry(description, args...)
}
