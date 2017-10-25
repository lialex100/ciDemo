package main

import (
	"github.com/gin-gonic/gin"
	"jwtService/handler"
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
