package main

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}

func main() {
	square := &square{side: 2}
	circle := &circle{radius: 3}

	areaCalculator := &areaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	fmt.Println()
}
