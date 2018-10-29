package controllers

import (
	"github.com/gin-gonic/gin"
	"jabargo/models"
	"time"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type UserController struct {
}

func (UserController) Get(context *gin.Context) {
	var mapParam = map[string]interface{}{}
	mapParam["key"] = context.Param("key")
	mapParam["paging"] = context.Param("paging")
	mapParam["count"] = context.Param("count")
	mapParam["limit"] = LimitPage
	mapParam["channelsId"] = []bson.ObjectId{"5bcdd71c9afd13895a508af8", "social"}
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
	if context.PostForm("password") != "" || context.PostForm("confirmPassword") != "" {
		if !strings.EqualFold(context.PostForm("password"), context.PostForm("confirmPassword")) {
			context.JSON(400, "password and confirm password must equal")
			return
		}
	}
	error, result := user.VerifyUser(context.PostForm("email"))
	if result != nil {
		context.JSON(400, "email already exist")
		return
	}
	if context.PostForm("email") == "" {
		context.JSON(400, "email must input")
		return
	}

	if context.PostForm("password") == "" {
		context.JSON(400, "password must input")
		return
	}
	user.Name = context.PostForm("name")
	user.Email = context.PostForm("email")
	user.Company = context.PostForm("company")
	user.Password = context.PostForm("password")
	user.Privilege = context.PostForm("privilege")
	user.Phone = context.PostForm("phone")
	user.Id = bson.NewObjectId()
	user.CreatedAt = time.Now()
	error, result = user.POST()
	if error != nil {
		context.JSON(500, error)
		return
	}
	context.JSON(200, &result)
}

func (UserController) Signin(context *gin.Context) {
	user := models.User{}
	error, result := user.VerifyUser(context.PostForm("email"))
	if error != nil || result == nil {
		context.JSON(404, "not found user or password")
		context.Abort()
	} else {
		error = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(context.PostForm("password")))
		if error != nil {
			context.JSON(404, "not found user or password")
			context.Abort()
			return
		}
		token := models.Token{}
		token.AccessToken = fmt.Sprint("Bearer ", bson.NewObjectId().Hex(), time.Now().UnixNano(), result.Id.Hex())
		token.UserId = result.Id
		token.User = models.UserChannel{bson.NewObjectId(), result.Name, result.Email, result.Company, result.Privilege, result.Phone}
		token.CreatedAt = time.Now()
		token.Create()
		context.JSON(200, &token)
	}
}

func (UserController) Update(context *gin.Context) {
	user := models.User{}
	if context.PostForm("password") != "" || context.PostForm("confirmPassword") != "" {
		if !strings.EqualFold(context.PostForm("password"), context.PostForm("confirmPassword")) {
			context.JSON(400, "password must equal")
			return
		}
	}
	user.Id = bson.ObjectId(context.PostForm("id"))
	user.Name = context.PostForm("name")
	user.Email = context.PostForm("email")
	user.Password = context.PostForm("password")
	user.Privilege = context.PostForm("privilege")
	user.Phone = context.PostForm("phone")
	user.CreatedAt = time.Now()
	error, result := user.POST()
	if error != nil {
		if strings.Contains(error.Error(), "duplicate key") {
			context.JSON(500, error)
		}
		return
	}
	context.JSON(200, &result)
}
