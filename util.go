package wgolib

import (
	"bytes"
	"encoding/json"
	"reflect"
)

// JSONEncode 将对象序列化为json str
// data is a struct, params: true:indent false:no_indent:default
func JSONEncode(data interface{}, params ...interface{}) string {
	jsons, errs := json.Marshal(data) //转换成JSON返回的是byte[]
	if errs != nil {
		return ""
	}
	if len(params) == 0 {
		return string(jsons)
	}

	if reflect.TypeOf(params[0]).Kind() == reflect.Bool && params[0].(bool) == false {
		return string(jsons)
	}
	var out bytes.Buffer
	err := json.Indent(&out, jsons, "", "  ")
	if err != nil {
		return string(jsons)
	}
	return out.String()
}
