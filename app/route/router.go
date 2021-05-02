package router

import (
	"BadOrange/app/service"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	route := gin.Default()

	route.GET("/", service.Root)

	route.POST("/signal/PushAll", service.PushAll)

	route.POST("/signal/Push", service.Push)

	return route

}