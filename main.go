package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GmAPI/controller"
	"github.com/GmAPI/model"

)

func main() {
	db, err := model.InitDB()
	if err != nil {
		fmt.Errorf("db: %v", err)
    }
	defer db.Close()
	router := gin.Default()
	router.Use(Cors())
	api :=router.Group("/api/g")
	{
		api.GET("/ping",controller.IndexGet)
		api.GET("/page/:id", controller.PageGet)
		api.GET("/all/page", controller.PageGetAll)
		api.POST("/add/page", controller.PageAdd)

		api.GET("/trainSearch/get/:id", controller.GetTrain)
		api.GET("/trainSearch/page", controller.GetTrainPage)
		api.POST("/trainSearch/add", controller.AddTrain)
		api.DELETE("/trainSearch/del/:id", controller.DelTrain)
		api.PUT("/trainSearch/edit", controller.EditTrain)
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}

func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        // c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        // if c.Request.Method == "OPTIONS" {
        //     c.AbortWithStatus(200)
        //     return
        // }
        c.Next()
    }
}