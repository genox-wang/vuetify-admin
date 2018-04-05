package controller

import (
	"net/http"
	"strconv"
	"vuetify-admin-api/app/model"

	"github.com/gin-gonic/gin"
)

// ServerAllGet  get all servers
func ServerAllGet(c *gin.Context) {
	if servers, err := new(model.Server).GetAll(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"servers": servers,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}

// ServerCreatePost create new server
func ServerCreatePost(c *gin.Context) {
	var server model.Server
	if err := c.BindJSON(&server); err == nil {
		if err := server.Create(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "create server success",
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

// ServerUpdatePut update server
func ServerUpdatePut(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "id <= 0",
		})
		return
	}

	var server model.Server
	if err := c.BindJSON(&server); err == nil {
		server.ID = uint(id)
		if err := server.Update(); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "update server success",
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

// ServerDelete delete server
func ServerDelete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "id <= 0",
		})
		return
	}

	server := &model.Server{}
	server.ID = uint(id)

	if err := server.Delete(); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "soft delete server success",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
}
