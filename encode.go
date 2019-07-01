package tag_marshall

import (
	"reflect"
)

var (
	keyCon   = "json"
	valueCon = "encode"
)

const (
	ignore = "-"
)

type MarshallRes map[string]interface{}

func SetKV(k, v string) {
	keyCon = k
	valueCon = v
}

func Marshall(v interface{}) MarshallRes {
	v2 := reflect.ValueOf(v)
	if !v2.IsValid() {
		return nil
	}
	return encode(v2)
}
