package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func New(addr string, handler *gin.Engine) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (s *Server) Serve() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewRouter() *gin.Engine {
	return gin.Default()
}
