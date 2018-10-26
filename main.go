package main

import (
		"os"
	"jabargo/models"
	"github.com/gin-gonic/gin"
	"jabargo/controllers"
)

func init() {
	models.GetSession();
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	v1 := router.Group("/v1/api")
	{
		v1.GET("/users", controllers.UserController{}.Get)
		v1.POST("/users", controllers.UserController{}.Post)
		v1.GET("/channels", controllers.ChannelController{}.Get)
		v1.POST("/channels", controllers.ChannelController{}.Post)
		v1.PUT("/channels/addUser", controllers.ChannelController{}.AddUser)
		v1.PUT("/channels/removeUser", controllers.ChannelController{}.RemoveUser)

	}
	router.Run(":" + port)
}
