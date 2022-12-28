package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/shaneikennedy/errorcodes/httpstatuscodes"
	"github.com/stretchr/testify/assert"
)

func TestCodeRoutes(t *testing.T) {
	// Arrange
	router := setupRouter()

	// Act Assert
	for _, code := range httpstatuscodes.StatusCodes {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/"+strconv.Itoa(code), nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, code, w.Code)

	}
}
