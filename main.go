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
	g := SetupGin()
	g.Run(":8080")

}

func SetupGin() *gin.Engine {
	g := gin.Default()
	g = GinRouter(g)
	return g
}

func GinRouter(engine *gin.Engine) *gin.Engine {
	engine.Static("/static", "public")
	engine.GET("/", handler.GetRoot)
	engine.GET("/productList", handler.ProductList)

	// return i am still alive
	engine.GET("/heartbeat", handler.HeartbeatHandler)
	return engine
}

func Add(a, b int) int {
	if a == 1 && b == 1 {
		return 3
	}
	return a + b
}
