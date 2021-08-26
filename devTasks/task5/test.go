package main

import (
	"testing"
)

func Test001(t *testing.T) {
	want := 2

	flaga := 1
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	got, _ := FindStr(ps, lookFor)
	Print_String_After(ps, got, flaga)
	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}

}

func Test002(t *testing.T) {
	want := 2
	flaga := 1
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	got, _ := FindStr(ps, lookFor)
	Print_String_Before(ps, got, flaga)
	if want != got {
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
	flagc := 2
	want := "No match found"
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "dfsdfsdfs"
	got, goterr := FindFullStr(ps, lookFor)
	Print_String_Around(ps, got, flagc)
	if want != goterr.Error() {
		t.Errorf("Got == %q, want %q", got, want)
	}
}

func Test005(t *testing.T) {
	flagc := 5
	want := "Error out of len"
	ps := []string{"The FlagSet", " type allows one", "to define independent", "sets of flags"}
	lookFor := "type"
	got, _ := FindFullStr(ps, lookFor)
	goterr := Print_String_Around(ps, got, flagc)
	if want != goterr.Error() {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
