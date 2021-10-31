package user

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {

    router.GET("/users", func(c *gin.Context) {
		c.JSON(200, GetUsers())
	})

	router.GET("/users/:id", func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, "Bad Parameter")
		}
		c.JSON(200, GetSingleUser(userId))
	})
}