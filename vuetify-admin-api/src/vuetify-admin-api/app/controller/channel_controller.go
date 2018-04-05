package controller

import (
	"net/http"
	"strconv"
	"vuetify-admin-api/app/model"

	"github.com/gin-gonic/gin"
)

// ChannelAllGet  get all channels
func ChannelAllGet(c *gin.Context) {
	if channels, err := new(model.Channel).GetAll(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"channels": channels,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

// ChannelCreatePost create channel
func ChannelCreatePost(c *gin.Context) {
	var channel model.Channel
	if err := c.BindJSON(&channel); err == nil {
		if err := channel.Create(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "create channel success",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

// ChannelUpdatePut update channel
func ChannelUpdatePut(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "id <= 0",
		})
		return
	}

	var channel model.Channel
	if err := c.BindJSON(&channel); err == nil {
		channel.ID = uint(id)
		if err := channel.Update(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "save channel success",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

// ChannelDelete delete channel
func ChannelDelete(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "id <= 0",
		})
		return
	}

	channel := &model.Channel{}
	channel.ID = uint(id)

	if err := channel.Delete(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "soft delete channel success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

// ChannelGetByToken get channel by token
func ChannelGetByToken(c *gin.Context) {
	token := c.Param("token")
	channel := &model.Channel{UUID: token}
	if err := channel.GetByToken(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	} else {
		servers := make([]string, 0)
		for _, v := range channel.Servers {
			servers = append(servers, v.Token)
		}
		c.JSON(http.StatusOK, gin.H{
			"name":    channel.Name,
			"servers": servers,
		})
	}
}
