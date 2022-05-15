package server

import (
	"kitchen/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/run-code", RunCodeHandler)
	router.Run("localhost:8080")
}

func RunCodeHandler(c *gin.Context) {
	cmd := c.Request.FormValue("cmd")
	input := c.Request.FormValue("input")

	result := lib.RunCode(cmd, input)
	c.String(http.StatusOK, result.String())
}
