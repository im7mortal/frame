package main

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
)


func TestContact(t *testing.T) {

	// RUN

	router.POST("/contact", handlerContact)
	req, err := http.NewRequest("POST", "/contact", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Success.\"}\n"))

	return
}

