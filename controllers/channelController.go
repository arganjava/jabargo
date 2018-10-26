package controllers

import (
	"github.com/gin-gonic/gin"
	"jabargo/models"
	"gopkg.in/mgo.v2/bson"
	"time"
)


type ChannelController struct {

}

func (ChannelController) Get(context *gin.Context) {
	var channel models.Channel
	error, result := models.Channel{}.Get(&channel)
	if error != nil {
		context.JSON(500, error)
	}
	context.JSON(200, &result)
}

func (ChannelController) Post(context *gin.Context) {
	channel := models.Channel{}
	channel.Name = context.PostForm("name")
	channel.Type = context.PostForm("type")
	channel.CreatedAt = time.Now()
	error, result := channel.Post()
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}

func (ChannelController) AddUser(context *gin.Context) {
	channel := models.Channel{}
	channel.Id = bson.ObjectIdHex(context.PostForm("channelId"))
	userId := bson.ObjectIdHex(context.PostForm("userId"))
	error, result := channel.AddUser(&userId)
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}

func (ChannelController) RemoveUser(context *gin.Context) {
	channel := models.Channel{}
	channel.Id = bson.ObjectIdHex(context.PostForm("channelId"))
	userId := bson.ObjectIdHex(context.PostForm("userId"))
	error, result := channel.RemoveUser(&userId)
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}