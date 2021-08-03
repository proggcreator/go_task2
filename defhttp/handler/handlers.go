package handler

import (
	"github.com/proggcreator/go_task2/defhttp/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}

}
