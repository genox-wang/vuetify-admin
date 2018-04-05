package controller

import (
	"net/http"
	"vuetify-admin-api/app/middleware"
	"vuetify-admin-api/app/model"

	"github.com/gin-gonic/gin"
)

// UserLoginPost is a function
func UserLoginPost(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err == nil {
		if err := user.Login(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"token": middleware.GetJWTToken(user),
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// UserCreatePost create user
func UserCreatePost(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err == nil {
		if err := user.Create(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "create user success",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "param error!",
	})
}

// UserAllGet get all user
func UserAllGet(c *gin.Context) {
	users := make([]*model.User, 0)
	if err := model.DB.Select("id, display_name, username").Find(&users).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":   err.Error(),
			"users": users,
		})
	}
}
