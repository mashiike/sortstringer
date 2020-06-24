package sortstringer

import (
	"fmt"
	"reflect"
	"sort"
)

var (
	//Ascending is a StringerSlice less function. sort order is ascending.
	Ascending = func(is, js string) bool { return is < js }
	//Descending is a StringerSlice less function. sort order is descending.
	Descending = func(is, js string) bool { return is > js }
)

//StringerSlice attaches the methods of sort.Interface that fmt.Stringer slice sorts.
type StringerSlice struct {
	rv      reflect.Value
	less    func(is, js string) bool
	swapper func(i, j int)
}

//NewStringerSlice is a constractor for *StringerSlice. sort order is Ascending
func NewStringerSlice(slice interface{}) *StringerSlice {
	return newStringerSlice(slice, Ascending)
}

func newStringerSlice(slice interface{}, less func(is, js string) bool) *StringerSlice {
	return &StringerSlice{
		rv:      reflect.ValueOf(slice),
		less:    less,
		swapper: reflect.Swapper(slice),
	}
}

// Len is the number of elements in the collection.
func (s *StringerSlice) Len() int {
	return s.rv.Len()
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s *StringerSlice) Less(i, j int) bool {
	iv, jv := s.rv.Index(i), s.rv.Index(j)
	return s.less(getString(iv), getString(jv))
}

// Swap swaps the elements with indexes i and j.
func (s *StringerSlice) Swap(i, j int) {
	s.swapper(i, j)
}

var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

// Slice sorts the provided slice given the provided less function.
// Same as sort.Slice function
func Slice(slice interface{}, less func(is, js string) bool) {
	s := newStringerSlice(slice, less)
	sort.Sort(s)
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
