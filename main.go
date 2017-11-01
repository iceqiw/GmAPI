package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GmApi/controller"
	"github.com/GmApi/model"

)

func main() {
	db, err := model.InitDB()
	if err != nil {
		fmt.Errorf("db: %v", err)
    }
	defer db.Close()
	router := gin.Default()
	api :=router.Group("/api/g")
	api.Use(Cors())
	{
		api.GET("/ping",controller.IndexGet)
		api.GET("/page/:id", controller.PageGet)
		api.GET("/all/page", controller.PageGetAll)
		api.POST("/add/page", controller.PageAdd)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// c.Writer.Header().Set("content-type ", "application/json; charset=utf-8")
        c.Next()
    }
}