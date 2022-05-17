package service

import (
	"kitchen/lib"
	"net/http"

	c "kitchen/config"

	"github.com/gin-gonic/gin"
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
	result, err := lib.OttoRunCode(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.String(http.StatusOK, result.String())
}
