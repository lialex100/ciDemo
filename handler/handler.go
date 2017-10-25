package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lialex100/ciDemo/helper"
	"log"
	"net/http"
)

const notImplement = "| Not Implement Yet |"

func GetRoot(c *gin.Context) {
	//c.String(http.StatusOK, "jwt service started")
	html := `<html>
				<head>
				</head>
				<body>
					<h1>
					Test 1
					</h1>
				</body>
			</html>`

	c.Data(200, "text/html; charset=utf-8", []byte(html))

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
