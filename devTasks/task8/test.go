package main

import (
	"errors"
	"testing"
)

func Test001(t *testing.T) {
	listargs := []string{}
	want := errors.New("Error: few netcat arguments")

	got := ClientNetcat(listargs)

	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test002(t *testing.T) {
	listargs := []string{"nc", "udp", "192.168.0.113", "80"}
	var want error
	got := ClientNetcat(listargs)
	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test003(t *testing.T) {
	listargs := "nc cd"
	want := errors.New("path required")

	got := execInput(listargs)

	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test004(t *testing.T) {
	listargs := []string{"nc", "tcp", "192.168.0.113", "80"}
	var want error
	got := ClientNetcat(listargs)
	if want != got {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test005(t *testing.T) {
	listargs := []string{"nc", "efsdf"}

	got := ClientNetcat(listargs)
	if got == nil {
		t.Errorf("Got == %q, want error", got)
	}
}
