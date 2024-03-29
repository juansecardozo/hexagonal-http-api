package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server/handler/courses"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server/handler/health"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server/middleware/logging"
	"github.com/juansecardozo/hexagonal-http-api/internal/platform/server/middleware/recovery"
	"github.com/juansecardozo/hexagonal-http-api/kit/command"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpPort string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	//deps
	commandBus command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpPort: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server is running on port", s.httpPort)

	srv := &http.Server{
		Addr:    s.httpPort,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.Use(recovery.Middleware(), logging.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
