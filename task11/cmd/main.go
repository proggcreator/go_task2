package main

import (
	"net/http"

	defhttp "github.com/proggcreator/go_task2/task11"
	myhand "github.com/proggcreator/go_task2/task11/handler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //формат для логгера json
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs")
	}

	//handler := http.HandlerFunc(defhttp.Serve)
	s := new(defhttp.Server)
	stor := myhand.NewStore()
	router := http.HandlerFunc(stor.Serve)

	if err := s.Run(viper.GetString("port"), router); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

	/*if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}*/
}
