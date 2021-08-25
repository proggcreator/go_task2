package main

import (
	"fmt"
	"testing"
)

func TestSortk001(t *testing.T) {
	ps := []string{"a f g h", "i f k l", " r e b y"}

	//want := [][]string{{}, {}}
	got := mySort(ps, 0, false, false, false)

	fmt.Println(got)

	//if got != want {
	//t.Errorf("ps.mySort() == %q, want %q", got, want)
	//}
}
