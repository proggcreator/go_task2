package main

import (
	"reflect"
	"testing"
)

func Test001(t *testing.T) {

	masstr := []string{"пятак", "тапок", "тяпка", "дом", "улица"}
	want := map[string]string{"пятак": "тяпка"}
	got := myanagrams(masstr)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test002(t *testing.T) {

	masstr := []string{"пятак", "тапок", "тяпка", "дом", "улица", "мод"}
	want := map[string]string{"пятак": "тяпка", "дом": "мод"}
	got := myanagrams(masstr)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test003(t *testing.T) {

	masstr := []string{"тапок", "тяпка", "дом", "улица"}
	want := map[string]string{}
	got := myanagrams(masstr)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test004(t *testing.T) {

	masstr := []string{"тапок"}
	want := map[string]string{}
	got := myanagrams(masstr)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
func Test005(t *testing.T) {

	masstr := []string{}
	want := map[string]string{}
	got := myanagrams(masstr)

	if reflect.DeepEqual(want, got) {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
