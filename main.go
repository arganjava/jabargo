package main

import (
	"os"
	"jabargo/models"
	"github.com/gin-gonic/gin"
	"jabargo/controllers"
	"fmt"
)

func init() {
	models.GetSession();
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	v1 := router.Group("/v1/api")
	{
		v1.Use(authAccess)
		v1.GET("/users", controllers.UserController{}.Get)
		v1.GET("/channels", controllers.ChannelController{}.Get)
		v1.POST("/channels", controllers.ChannelController{}.Post)
		v1.PUT("/channels/addUser", controllers.ChannelController{}.AddUser)
		v1.PUT("/channels/removeUser", controllers.ChannelController{}.RemoveUser)

	}
	v1Auth := router.Group("/v1/auth")
	{
		v1Auth.POST("/signin", controllers.UserController{}.Signin)
		v1Auth.POST("/signup", controllers.UserController{}.Post)
	}

	router.Run(":" + port)
}

func authAccess(context *gin.Context) {
	if context.Request.Header["Authorization"] == nil {
		context.JSON(403, "unauthorized")
		context.Abort()
		return
	} else {
		accessToken := context.Request.Header["Authorization"][0]
		fmt.Print(context.Get(accessToken))
		_, tokenInfo := models.Token{}.FindToken(accessToken)
		if tokenInfo == nil {
			context.JSON(403, "invalid token")
			context.Abort()
			return
		}
	}
	context.Next()
}
