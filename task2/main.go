package main

import (
	"fmt"
	"log"
)

func main() {

	var s string

	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(UnpacStr(s))
}
