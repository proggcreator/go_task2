package main

import "fmt"

func a() {

	fmt.Println("a")

}
func b(x int) {
	fmt.Println("b:", x)
}

func main() {
	x := 5
	defer b(x)
	a()
	x = 4
	return

}
