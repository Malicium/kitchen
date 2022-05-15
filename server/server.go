package server

import (
	"kitchen/lib"
	"net/http"

	"github.com/gin-gonic/gin"
	c "kitchen/config"
)

var config = c.Get()

func Start() {
	router := gin.Default()
	router.SetTrustedProxies([]string{config.TrustedProxy})

	router.POST("/run-code", RunCodeHandler)

	router.Run(config.HttpAddress)
}

func RunCodeHandler(c *gin.Context) {
	input := c.Request.FormValue("input")
	result, err := lib.RunCode(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.String(http.StatusOK, result.String())
}
