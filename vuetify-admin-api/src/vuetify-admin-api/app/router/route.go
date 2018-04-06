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
		authorized.POST("/user/:id", controller.UserUpdatePost)
		authorized.DELETE("/user/:id", controller.UserDelete)
	}
}
