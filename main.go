package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lialex100/ciDemo/handler"
)

func main() {
	go StartGin()
	select {}
}

func StartGin() {
	g := gin.Default()
	g.Static("/static", "public")
	GinRouter(g)
	g.Run(":8080")
}

func GinRouter(engine *gin.Engine) {

	engine.GET("/", handler.GetRoot)
	engine.GET("/productList", handler.ProductList)
	// return i am still alive
	engine.GET("/heartbeat", handler.HeartbeatHandler)
}

func Add(a, b int) int {
	if a == 1 && b == 1 {
		return 3
	}
	return a + b
}
