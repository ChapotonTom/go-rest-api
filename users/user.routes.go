package user

import (
	"restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
    router.GET("/users", middleware.AuthorizeEmployee, HandleGetUsers)
	router.GET("/users/:id", middleware.AuthorizeUser, HandleGetUser)
	router.POST("/users", middleware.AuthorizeManager, HandleUserCreate)
	router.PUT("/users/:id", middleware.AuthorizeManager, HandleUserUpdate)
}