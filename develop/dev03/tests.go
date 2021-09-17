package main

import (
	"fmt"
	"testing"
)

func TestSortk001(t *testing.T) {
	ps := []string{"a f g h", "i f k l", " r e b y"}

	want := [][]string{{"a", "f", "g", "h"}, {"i", "f", "k", "l"}, {"r", "e", "b", "y"}}

	got := mySort(ps, 0, false, false, false)

	fmt.Println(got)
	for i, _ := range want {
		for j, _ := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("ps.mySort() == %q, want %q", got, want)
			}
		}
	}
}

func TestSortk002(t *testing.T) {
	ps := []string{"asd rft edc", "wsv ret rev", "tgy rev trb"}

	want := [][]string{{"wsv", "ret", "rev", "tgy", "rev", "trb", "asd", "rft", "edc"}}

	got := mySort(ps, 1, false, true, true)

	fmt.Println(got)
	for i, _ := range want {
		for j, _ := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("ps.mySort() == %q, want %q", got, want)
			}
		}
	}
}
func TestSortk003(t *testing.T) {
	ps := []string{"1 2 3 4 ", "5 6 7 8", "9 10 11 12"}

	want := [][]string{{"9", "10", "11", "12"}, {"5", "6", "7", "8"}, {"1", "2", "3", "4"}}

	got := mySort(ps, 0, true, true, false)

	fmt.Println(got)
	for i, _ := range want {
		for j, _ := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("ps.mySort() == %q, want %q", got, want)
			}
		}
	}
}
func TestSortk004(t *testing.T) {
	ps := []string{"asd rft edc", "tgy rev trb", "wsv ret rev"}
	want := [][]string{{"asd", "rft", "edc"}, {"wsv", "ret", "rev"}, {"tgy", "rev", "trb"}}
	got := mySort(ps, 0, false, false, false)

	fmt.Println(got)
	for i, _ := range want {
		for j, _ := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("ps.mySort() == %q, want %q", got, want)
			}
		}
	}
}
func TestSortk005(t *testing.T) {
	ps := []string{"1 2 3 ", "4 5 6 ", "1 2 3"}

	want := [][]string{{"4", "5", "6"}, {"1", "2", "3"}}

	got := mySort(ps, 0, false, true, true)

	fmt.Println(got)
	for i, _ := range want {
		for j, _ := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf("ps.mySort() == %q, want %q", got, want)
			}
		}
	}
}
