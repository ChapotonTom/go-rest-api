package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine) {
    router.POST("/login", HandleLogin)
}