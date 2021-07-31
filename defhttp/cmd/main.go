package main

import (
	"fmt"
	"net/http"

	"github.com/proggcreator/go_task2/defhttp/handler"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter)) //формат для логгера json
	//if err := initConfig(); err != nil {
	//	logrus.Fatalf("error initializing configs")
	//}
	//иерархия зависимостей
	//services := service.NewService()
	//handlers := handler.NewHandler(services)

	mux := http.NewServeMux()
	mux.Handle("/", handler.Homee)
	//mux.Handle("/create_event/",)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}

/*
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

*/
