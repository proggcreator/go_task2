package main

type customError struct {
	msg string
}

func (e *customError) Error() string { //метод типа customError
	return e.msg //возвращает сообщение
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil { //nil!=nil
		println("error")
		return
	}
	println("ok")
}

//error
