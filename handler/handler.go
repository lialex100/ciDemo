package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lialex100/ciDemo/helper"
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

func HeartbeatHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
