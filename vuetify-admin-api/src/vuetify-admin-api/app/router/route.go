package router

import (
	// "net/http"
	// "vuetify-admin-api/app/config"
	"vuetify-admin-api/app/config"
	"vuetify-admin-api/app/controller"
	"vuetify-admin-api/app/middleware"
	// "github.com/gin-gonic/gin"
)

func route() {
	router.POST("/console/login", controller.UserLoginPost)

	authorized := router.Group("/console")

	if config.GetBool("authRequired") {
		authorized.Use(middleware.JWTAuthRequired())
	}

	{
		authorized.POST("/user/", controller.UserCreatePost)
		authorized.GET("/user/all", controller.UserAllGet)

		authorized.GET("/channel/all", controller.ChannelAllGet)
		authorized.POST("/channel/", controller.ChannelCreatePost)
		authorized.PUT("/channel/:id", controller.ChannelUpdatePut)
		authorized.DELETE("/channel/:id", controller.ChannelDelete)

		authorized.GET("/server/all", controller.ServerAllGet)
		authorized.POST("/server/", controller.ServerCreatePost)
		authorized.PUT("/server/:id", controller.ServerUpdatePut)
		authorized.DELETE("/server/:id", controller.ServerDelete)
	}

	api := router.Group("/api")
	{
		api.GET("/channel/:token", controller.ChannelGetByToken)
	}
}
