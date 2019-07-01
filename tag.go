package tag_marshall

import (
	"fmt"
	"reflect"
)

func encode(v reflect.Value) map[string]interface{} {

	if !v.IsValid() {
		return nil
	}

	t := v.Type()

	switch t.Kind() {
	case reflect.Ptr:
		return encode(v.Elem())

	case reflect.Interface:
		return encode(v.Elem())

	case reflect.Slice:
		if v.Len() == 0 {
			return nil
		}
		return encode(v.Index(0))

	case reflect.Struct:
		t := v.Type()
		tags := make(MarshallRes)
		for i := 0; i < t.NumField(); i++ {

			key := t.Field(i).Tag.Get(keyCon)
			if key == ignore {
				continue
			}
			value := t.Field(i).Tag.Get(valueCon)
			if value == ignore {
				continue
			}

			tags[key] = value

			if k := t.Field(i).Type.Kind(); k == reflect.Ptr ||
				k == reflect.Slice ||
				k == reflect.Struct ||
				k == reflect.Interface {

				tags[fmt.Sprintf("%s_struct", key)] = encode(v.Field(i))
			}
		}
		return tags
	}

	return nil
}
