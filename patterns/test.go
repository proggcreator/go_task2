package main

import (
	"fmt"
)

type tree struct {
	next  *tree
	value int
}

func prit(tr *tree) {

	for tr != nil {

		fmt.Println(tr.value)
		if tr.next != nil {
			fmt.Println(tr)
		}
		tr = tr.next

	}
}

func main() {

	mytree1 := new(tree)
	mytree2 := new(tree)

	mytree1.value = 5
	mytree1.next = mytree2
	mytree2.value = 10

	prit(mytree1)

}
