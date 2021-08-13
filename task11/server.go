package defhttp

import (
	"net/http"
)

type Handler struct {
	h *http.Handler
}

func NewHandler() *Handler {
	return &Handler{}
}

type Server struct {
	httpServer *http.Server
}
