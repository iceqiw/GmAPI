package controller

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/GmApi/model"
	"github.com/GmApi/helper"
)

func GetTrain(c *gin.Context) {
	id, _ := helper.StrTo(c.Param("id")).Int64()
	TrainSearchModel := &model.TrainSearch{Id:id}
	trainSearch, has := TrainSearchModel.Get()
	if !has {
		c.JSON(200, trainSearch)
		return
	}
}

func GetTrainPage(c *gin.Context) {
	TrainSearchModel := &model.TrainSearch{}
	trainSearchs, err := TrainSearchModel.GetList(1,5)
	if err == nil {
		c.JSON(200, trainSearchs)
		return
	}
	log.Println("Println err",err)
}

func AddTrain(c *gin.Context) {
	TrainSearchModel := model.TrainSearch{}
	if c.BindJSON(&TrainSearchModel) == nil {
		ok,err:=TrainSearchModel.Insert()
		log.Println("Println err",err)
		c.JSON(200, gin.H{
			"sussces": ok,
		})
	}
}

