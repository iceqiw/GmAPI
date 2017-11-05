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
	trainSearchs, err := TrainSearchModel.All()
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
		if err==nil{
			c.JSON(200, gin.H{
				"sussces": ok,
			})
			return
		}
		log.Println("Println err",err)	
	}
}

func DelTrain(c *gin.Context) {
	id, _ := helper.StrTo(c.Param("id")).Int64()
	TrainSearchModel := &model.TrainSearch{Id:id}
	ok, err := TrainSearchModel.Delete()
	if err==nil{
		c.JSON(200, gin.H{
			"sussces": ok,
		})
		return
	}
}

func EditTrain(c *gin.Context) {
	TrainSearchModel := model.TrainSearch{}
	if c.BindJSON(&TrainSearchModel) == nil {
		ok,err:=TrainSearchModel.Update()
		if err==nil{
			c.JSON(200, gin.H{
				"sussces": ok,
			})
			return
		}
		log.Println("Println err",err)	
	}
}