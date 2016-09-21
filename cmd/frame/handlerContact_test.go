package main

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"github.com/stretchr/testify/assert"
	"net/url"
	"bytes"
)


/*
func TestContact(t *testing.T) {

	// RUN
	router.POST("/contact", handlerContact)
	req, err := http.NewRequest("POST", "/contact", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte("{\"message\":\"Success.\"}\n"))

	return
}
*/


func TestContactValidation(t *testing.T) {
	router.POST("/contact", handlerContact)
	body := url.Values{}

	// RUN
	req, err := http.NewRequest("POST", "/contact", bytes.NewBufferString(body.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte(`{"statusCode":400,"error":"Bad Request","message":"child \"name\" fails because [\"name\" is required]","validation":{"source":"payload","keys":["name"]}}`))

	body.Set("name", "name")
	resp = getPostResponse(&body)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte(`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" is required]","validation":{"source":"payload","keys":["email"]}}`))

	body.Set("email", "invalidEmail")
	resp = getPostResponse(&body)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte(`{"statusCode":400,"error":"Bad Request","message":"child \"email\" fails because [\"email\" must be a valid email]","validation":{"source":"payload","keys":["email"]}}`))

	body.Set("email", "valid@email.com")
	resp = getPostResponse(&body)
	assert.Equal(t, 400, resp.Code)
	assert.Equal(t, resp.Body.Bytes(), []byte(`{"statusCode":400,"error":"Bad Request","message":"child \"message\" fails because [\"message\" is required]","validation":{"source":"payload","keys":["message"]}}`))
	return
}

func getPostResponse(body *url.Values) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", "/contact", bytes.NewBufferString(body.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	return resp
}


