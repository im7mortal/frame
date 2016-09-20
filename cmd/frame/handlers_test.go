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

	// RUN
	r.GET("/", handlerIndex)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Welcome to the plot device.\"}\n"))

	r.POST("/contact", handlerContact)
	req, err = http.NewRequest("POST", "/contact", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Success.\"}\n"))

	return
}

