package main

import "fmt"

func main() {
	builder1 := getBuilder("house1")
	builder2 := getBuilder("house2")

	director := newDirector(builder1)
	house1 := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", house1.doorType)
	fmt.Printf("Normal House Window Type: %s\n", house1.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", house1.floor)

	director.setBuilder(builder2)
	house2 := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", house2.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", house2.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", house2.floor)

}
