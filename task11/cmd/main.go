package main

import (
	"log"
	"net/http"

	defhttp "github.com/proggcreator/go_task2/task11"
)

func main() {

	router := http.HandlerFunc(defhttp.ServeHTTP)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
