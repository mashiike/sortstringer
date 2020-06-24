package sortstringer

import (
	"fmt"
	"reflect"
	"sort"
)

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

func Slice(slice interface{}, less func(is, js string) bool) {
	rv := reflect.ValueOf(slice)
	if rv.Len() == 0 {
		return
	}
	sort.Slice(slice, func(i, j int) bool {
		iv, jv := rv.Index(i), rv.Index(j)
		return less(getString(iv), getString(jv))
	})
}

func getString(rv reflect.Value) string {
	if rv.Kind() == reflect.String {
		return rv.String()
	}
	if rv.Type().Implements(stringerType) {
		return rv.Interface().(fmt.Stringer).String()
	}
	if rv.Kind() == reflect.Interface || rv.Kind() == reflect.Ptr {
		return getString(rv.Elem())
	}
	return ""
}
