package main

func main() {
	ch := make(chan int) //не буфферизированный
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //в канал
		}
		//close(ch)
	}()

	for n := range ch { //считывает из канала
		println(n)
	}
	//когда считает все - додлок т.к ожидание бесконечное
}
