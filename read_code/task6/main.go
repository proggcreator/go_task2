package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s) //у прежнего слайса не изменился размер
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"         //меняем 0ой элемент
	i = append(i, "4") //создание расширенного слайса
	i[1] = "5"         //меняем 1ый элемент
	i = append(i, "6") //создание расширенного слайса
	fmt.Println(i)
}

//3 2 3
