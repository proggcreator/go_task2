package main

import (
	"testing"
)

func TestUnpack001(t *testing.T) {
	var ps string
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = UnpacStr(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
func TestUnpack002(t *testing.T) {
	var ps string
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = UnpacStr(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
