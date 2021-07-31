package main

import (
	"log"
	"net/http"

	"github.com/proggcreator/go_task2/defhttp"
)

func main() {

	router := http.HandlerFunc(defhttp.Serve)

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
