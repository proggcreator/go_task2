package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //формат для логгера json
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs")
	}

	//иерархия зависимостей
	//services := service.NewService()
	//handlers := handler.NewHandler(services)

	http.Handle("/", handler.homeeee)
	//http.Handle("/about", handler.main)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
