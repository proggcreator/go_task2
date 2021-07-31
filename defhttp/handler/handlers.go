package handler

import (
	"fmt"

	"github.com/proggcreator/go_task2/defhttp/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}

}

func (handle *Handler) Homee() {
	fmt.Print("Heeeeeeeeeeeeeeeeeeloo!")

}
