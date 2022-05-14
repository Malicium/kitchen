package main

import (
	"kitchen/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/code", func(c *gin.Context) {
		lib.RunCode("node", "hello.js")
	})
	router.Run("localhost:8080")
}
