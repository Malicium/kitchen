package server

import (
	"kitchen/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})

	router.POST("/run-code", RunCodeHandler)

	router.Run("localhost:8080")
}

func RunCodeHandler(c *gin.Context) {
	cmd := c.Request.FormValue("cmd")
	input := c.Request.FormValue("input")

	result, err := lib.RunCode(cmd, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.String(http.StatusOK, result.String())
}
