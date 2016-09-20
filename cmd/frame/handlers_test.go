package main

import (
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"os"
	"flag"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	flag.Parse()

	// SETUP
	gin.SetMode(gin.TestMode)
	router = gin.Default()

	os.Exit(m.Run())
}





func TestIndex(t *testing.T) {

	// RUN
	router.GET("/", handlerIndex)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Welcome to the plot device.\"}\n"))

	router.POST("/contact", handlerContact)
	req, err = http.NewRequest("POST", "/contact", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Success.\"}\n"))

	return
}

