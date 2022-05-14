package server

import (
	"kitchen/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.2"})
	GetRoutes(router)
	router.Run("localhost:8080")
}

func GetRoutes(router *gin.Engine) {
	router.GET("/run-code/:mode/:file", RunCode)
}

func RunCode(c *gin.Context) {
	mode := c.Param("mode")
	file := c.Param("file")

	output := lib.RunCode(mode, file)
	c.String(http.StatusOK, output.String())
}
