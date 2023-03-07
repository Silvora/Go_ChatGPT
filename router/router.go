package router

import (
	"Go_ChatGPT/service"
	"Go_ChatGPT/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	route := gin.Default()
	route.Use(utils.Cors())

	route.POST("/ai", service.Chat)
	route.GET("/help", service.Help)

	return route
}
