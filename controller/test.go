package controller

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/GmApi/model"
	"github.com/GmApi/helper"
)


func IndexGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func PageGet(c *gin.Context) {
	id, _ := helper.StrTo(c.Param("id")).Int64()
	userModel := &model.User{Id:id}
	user, has := userModel.Get()
	if !has {
		c.JSON(200, gin.H{
			"message": user,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user",
	})
}

func PageGetAll(c *gin.Context) {
	userModel := &model.User{}
	users, _ := userModel.GetList(1,5)
	c.JSON(200, users)
}

