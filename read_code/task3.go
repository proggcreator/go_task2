package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil //указатель не указывает ни на что
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        //nil
	fmt.Println(err == nil) //false (nil!=nil )
}

// интерфейс определяет и описывает конкретные методы, которые должны быть у какого-то другого типа.
//любой объект удовлетворяет пустому интерфейсу (map[string]interface{})
