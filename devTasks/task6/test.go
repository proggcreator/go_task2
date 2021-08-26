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
	separator := ","
	strlist := []string{"a,d,f,g", "e,r,t,y", "f,v,b,n"}
	flagf := "0,2"
	flags := true

	want := []int{1, 2, 3, 4, 5}
	got := myCut(strlist, separator, flagf, flags)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test005(t *testing.T) {
	separator := "   "
	strlist := []string{"1   2   3", "4   5   6", "7   8   9"}
	flagf := "100"
	flags := true

	want := "Error out of len"
	_, got := myCut(strlist, separator, flagf, flags)

	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
