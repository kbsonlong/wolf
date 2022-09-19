package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kbsonlong/wolf/internal/service"
)

func InitRouter() gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")

	api := r.Group("/api")
	{
		api.GET("/ping", service.Ping)
		api.POST("/events", service.SpotEvent)
	}

	return *r
}
