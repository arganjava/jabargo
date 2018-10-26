package controllers

import (
	"github.com/gin-gonic/gin"
	"jabargo/models"
	"time"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {

}

func (UserController) Get(context *gin.Context) {
	var mapParam = map[string]interface{}{}
	mapParam["key"] = context.Param("key")
	mapParam["paging"] = context.Param("paging")
	mapParam["count"] = context.Param("count")
	mapParam["limit"] = LimitPage
	mapParam["channelsId"] = []bson.ObjectId{"5bcdd71c9afd13895a508af8","social"}
	var user models.User
	error, result := user.GET(&mapParam)
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}

func (UserController) Post(context *gin.Context) {
	user := models.User{}
	user.Name = context.PostForm("name")
	user.Email = context.PostForm("email")
	user.Company = context.PostForm("company")
	user.Password = context.PostForm("password")
	user.Privilege = context.PostForm("privilege")
	user.Phone = context.PostForm("phone")
	user.CreatedAt = time.Now()
	error, result := user.POST()
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}
