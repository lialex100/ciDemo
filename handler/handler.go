package handler

import (
	"github.com/gin-gonic/gin"
	"jwtService/helper"
	"log"
	"net/http"
)

const notImplement = "| Not Implement Yet |"

func GetRoot(c *gin.Context) {
	c.String(http.StatusOK, "jwt service started")
}

func ProductList(c *gin.Context) {
	log.Println(helper.GetFunctionName(ProductList))
	c.JSON(http.StatusOK, gin.H{"products": []string{
		"ApplePie",
		"BananaPie",
	}})
}
