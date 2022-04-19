package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const httpPort = ":8080"

func main() {
	println("Server running on", httpPort)

	srv := gin.New()
	srv.GET("/health", healthHandler)

	log.Fatal(srv.Run(httpPort))
}

func healthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Everything is OK!")
}