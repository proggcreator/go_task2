package main

import (
	"fmt"
	"log"

	u "github.com/proggcreator/go_task2/unpack"
)

func main() {

	var s string

	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(u.UnpacStr(s))
}
