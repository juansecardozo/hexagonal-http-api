package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server/handler/health"
	"log"
)

type Server struct {
	httpPort string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpPort: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server is running on port", s.httpPort)
	return s.engine.Run(s.httpPort)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
}
