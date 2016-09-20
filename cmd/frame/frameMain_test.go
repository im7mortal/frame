package main

import (
	"testing"
	"github.com/gin-gonic/gin"
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
