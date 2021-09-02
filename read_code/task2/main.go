package main

import "fmt"

//defer  вызывает перед завершением функции, после return
func test() (x int) { //возвращаем x
	defer func() { //в стек вызовов
		x++

	}()
	x = 1
	return // 2
}
func anotherTest() int { //возвращаем 1 потом defer
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x //1
}

func main() {
	fmt.Println(test())        //2
	fmt.Println(anotherTest()) //1
}
