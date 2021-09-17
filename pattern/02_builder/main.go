package main

import "fmt"

func main() {
	builder1 := getBuilder("house1")
	builder2 := getBuilder("house2")

	director := newDirector(builder1)
	house1 := director.buildHouse()

	fmt.Printf("Первый дом, крыша: %s\n", house1.roof)
	fmt.Printf("Первый дом, стены: %s\n", house1.walls)

	director.setBuilder(builder2)
	house2 := director.buildHouse()

	fmt.Printf("\nВторой дом, крыша: %s\n", house2.roof)
	fmt.Printf("Второй дом, стены:: %s\n", house2.walls)

}

/*Первый дом, крыша: Дерево
Первый дом, стены: Дерево

Второй дом, крыша: Стекло
Второй дом, стены:: Бетон*/
