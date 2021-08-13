package main

import (
	defhttp "github.com/proggcreator/go_task2/task11"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	router := http.HandlerFunc(defhttp.Server)

	srv := new(defhttp.Server)
	if err := srv.Run(viper.GetString("port"), router); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
