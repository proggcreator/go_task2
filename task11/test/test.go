package main

import (
	"fmt"
	"time"
)

func main() {
	mystr := "2007-01-02"
	myDate, err := time.Parse("2006-01-02", mystr)
	if err != nil {
		fmt.Println("Error")
	}

}
