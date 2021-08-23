package main

import (
	"testing"
)

func TestUnpack001(t *testing.T) {
	var ps string
	var got, want string

	ps = "a4bc2d5e"
	want = "aaaabccddddde"
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
func TestUnpack002(t *testing.T) {
	var ps string
	var got, want string

	ps = "asdf"
	want = "asdf"
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
func TestUnpack003(t *testing.T) {
	var ps string
	var got, want string

	ps = `qwe\\5`
	want = `qwe\\\\\`
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
func TestUnpack004(t *testing.T) {
	var ps string
	var got, want string

	ps = `qwe\13`
	want = `qwe111`
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}

func TestUnpack005(t *testing.T) {
	var ps string
	var got, want string

	ps = "123"
	want = ""
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
func TestUnpack006(t *testing.T) {
	var ps string
	var got, want string

	ps = "asd7231231"
	want = "asddddddd"
	got = Unpack(ps)

	if got != want {
		t.Errorf("ps.Unpack() == %q, want %q", got, want)
	}
}
