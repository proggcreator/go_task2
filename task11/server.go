package defhttp

import (
	"context"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Shutdown(ctx context.Context) error { //выход из приложения
	return s.httpServer.Shutdown(ctx)
}
