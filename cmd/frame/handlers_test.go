package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	// SETUP
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/", handlerIndex)

	// RUN
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Welcome to the plot device.\"}\n"))

	return
}

