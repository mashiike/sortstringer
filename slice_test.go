package sortstringer_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mashiike/sortstringer"
)

type MyString string

type Data struct{ Value int }

func newDataSlice(values []int) []Data {
	datas := make([]Data, 0, len(values))
	for _, v := range values {
		datas = append(datas, Data{Value: v})
	}
	return datas
}
func (d Data) String() string {
	return fmt.Sprintf("value=%d", d)
}

func TestSlice(t *testing.T) {
	pd := &Data{Value: 50}

	cases := []struct {
		casename string
		source   interface{}
		expected interface{}
	}{
		{
			casename: "reflect.String",
			source:   []MyString{"ddd", "ccc", "aaa", "bbb", "abc"},
			expected: []MyString{"aaa", "abc", "bbb", "ccc", "ddd"},
		},
		{
			casename: "fmt.Stringer",
			source:   newDataSlice([]int{2, 3, 1, 5}),
			expected: newDataSlice([]int{1, 2, 3, 5}),
		},
		{
			casename: "interface{}",
			source:   []interface{}{"a", 1, "b", pd, "c", "d"},
			expected: []interface{}{1, "a", "b", "c", "d", pd},
		},
		{
			casename: "no stringable",
			source:   []int{2, 3, 4, 5, 1},
			expected: []int{2, 3, 4, 5, 1},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case.%d %s", i+1, c.casename), func(t *testing.T) {
			t.Logf("before: %+v", c.source)
			sortstringer.Slice(c.source, func(is, js string) bool {
				t.Logf("is, js := %s, %s", is, js)
				return is < js
			})
			t.Logf("after: %+v", c.source)
			if !reflect.DeepEqual(c.source, c.expected) {
				t.Logf("actual:   %+v", c.source)
				t.Logf("expected: %+v", c.expected)
				t.Error("unexpected after source state")
			}
		})
	}
}

func TestSliceStable(t *testing.T) {
	pd := &Data{Value: 50}

	cases := []struct {
		casename string
		source   interface{}
		expected interface{}
	}{
		{
			casename: "reflect.String",
			source:   []MyString{"ddd", "dcd", "abc", "bbb", "aaa"},
			expected: []MyString{"abc", "aaa", "bbb", "ddd", "dcd"},
		},
		{
			casename: "fmt.Stringer",
			source:   newDataSlice([]int{2, 3, 1, 5}),
			expected: newDataSlice([]int{2, 3, 1, 5}),
		},
		{
			casename: "interface{}",
			source:   []interface{}{"ab", "aa", pd, "c", "d"},
			expected: []interface{}{"ab", "aa", "c", "d", pd},
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case.%d %s", i+1, c.casename), func(t *testing.T) {
			t.Logf("before: %+v", c.source)
			sortstringer.SliceStable(c.source, func(is, js string) bool {
				t.Logf("is, js := %s, %s", is, js)
				return is[0] < js[0]
			})
			t.Logf("after: %+v", c.source)
			if !reflect.DeepEqual(c.source, c.expected) {
				t.Logf("actual:   %+v", c.source)
				t.Logf("expected: %+v", c.expected)
				t.Error("unexpected after source state")
			}
		})
	}
}

func TestSliceIsSorted(t *testing.T) {
	cases := []struct {
		casename string
		slice    interface{}
		less     func(is, js string) bool
		expected bool
	}{
		{
			casename: "reflect.String Ascending",
			slice:    []MyString{"aaa", "abc", "adb", "cbd", "ccc"},
			less:     sortstringer.Ascending,
			expected: true,
		},
		{
			casename: "reflect.String not Descending",
			slice:    []MyString{"aaa", "abc", "adb", "cbd", "ccc"},
			less:     sortstringer.Descending,
			expected: false,
		},
		{
			casename: "fmt.Stringer Descending",
			slice:    newDataSlice([]int{5, 3, 2, 1}),
			less:     sortstringer.Descending,
			expected: true,
		},
		{
			casename: "fmt.Stringer not Ascending",
			slice:    newDataSlice([]int{5, 3, 2, 1}),
			less:     sortstringer.Ascending,
			expected: false,
		},
		{
			casename: "interface{} Ascending",
			slice:    []interface{}{1, "a", "b", "c", Data{Value: 2}, "z"},
			less:     sortstringer.Ascending,
			expected: true,
		},
		{
			casename: "interface{} not Descending",
			slice:    []interface{}{1, "a", "b", "c", Data{Value: 2}, "z"},
			less:     sortstringer.Descending,
			expected: false,
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case.%d %s", i+1, c.casename), func(t *testing.T) {
			t.Logf("slice: %+v", c.slice)
			actual := sortstringer.SliceIsSorted(c.slice, c.less)
			if actual != c.expected {
				t.Logf("actual:   %+v", actual)
				t.Logf("expected: %+v", c.expected)
				t.Error("unexpected is sorted state")
			}
		})
	}
}
