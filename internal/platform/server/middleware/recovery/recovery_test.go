package recovery

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecoveryMiddleware(t *testing.T) {
	// Setting up Gin server
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(Middleware())
	engine.GET("/test-middleware", func(context *gin.Context) {
		panic("unexpected behaviour")
	})

	// Setting up HTTP recorder and request
	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	// Asserting request does not produce a panic
	assert.NotPanics(t, func() {
		engine.ServeHTTP(httpRecorder, req)
	})
}
