package handler

import "net/http"

type Handler struct {
	handl *http.Handler
	contx *MyContext
}

func NewHandler() *Handler {
	return &Handler{
		handl: new(http.Handler),
		contx: NewMyContext(),
	}
}
