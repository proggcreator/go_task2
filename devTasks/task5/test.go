package main

import (
	"reflect"
	"testing"
)

func Test001(t *testing.T) {
	want := []string{" type allows one", "to define independent"}
	flaga := 1
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	inx, _ := FindStr(ps, lookFor)
	got, _ := Find_String_After(ps, inx, flaga)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}

}

func Test002(t *testing.T) {
	want := []string{"The FlagSet", " type allows one"}
	flagb := 1
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	inx, _ := FindStr(ps, lookFor)
	got, _ := Find_String_Before(ps, inx, flagb)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test003(t *testing.T) {
	want := 3
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "sets of flags"
	got, _ := FindFullStr(ps, lookFor)
	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test004(t *testing.T) {
	flagC := 2
	want := "No match found"
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "dfsdfsdfs"
	indx, _ := FindFullStr(ps, lookFor)
	_, got := Find_String_Around(ps, indx, flagC)
	if want != got.Error() {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test005(t *testing.T) {
	flagc := 5
	want := "Out range"
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	indx, _ := FindFullStr(ps, lookFor)
	_, got := Find_String_Around(ps, indx, flagc)
	if want != got.Error() {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
