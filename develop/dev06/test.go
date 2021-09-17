package main

import (
	"reflect"
	"testing"
)

func Test001(t *testing.T) {

	strlist := []string{"1   2   3", "4   5   6", "7   8   9"}
	separator := "   "
	want := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	got := CrmMatrix(strlist, separator)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test002(t *testing.T) {

	strlist := []string{"1|2|3", "4|5|6", "7|8|9"}
	separator := "|"
	want := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	got := CrmMatrix(strlist, separator)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test003(t *testing.T) {

	strlist := "1,2,3,4,5"

	want := []int{1, 2, 3, 4, 5}
	got := StrFflagToIntList(strlist)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test004(t *testing.T) {

	strlist := [][]string{{"i", "want", "to"}, {"gain "}, {"a", "order"}}
	want := [][]string{{"i", "want", "to"}, {"a", "order"}}
	got := FindOnlySep(strlist)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test005(t *testing.T) {

	strlist := [][]string{{}}
	want := [][]string{{}}
	got := FindOnlySep(strlist)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
