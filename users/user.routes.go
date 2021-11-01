package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
    router.GET("/users", HandleGetUsers)
	router.GET("/users/:id", HandleGetUser)
	router.POST("/users", HandleUserCreate)
	router.PUT("/users/:id", HandleUserUpdate)
}