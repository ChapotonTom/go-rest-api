package user

import (
	"strconv"
	"errors"
	"github.com/gin-gonic/gin"
)

func validateRoles(roles []string) error {
	for _, role := range roles {
		if role != "employee" && role != "manager" {
			return errors.New("Wrong role type")
		}
	}
	return nil
}

func UserRouter(router *gin.Engine) {

    router.GET("/users", func(c *gin.Context) {
		c.JSON(200, GetUsers())
	})

	router.GET("/users/:id", func(c *gin.Context) {
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Bad parameters"})
		}
		c.JSON(200, GetSingleUser(userId))
	})

	router.POST("/users", func(c *gin.Context) {
		var newUser User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(400, gin.H{"error": "Bad parameters"})
			return
		}
		if err := validateRoles(newUser.Roles); err != nil {
			c.JSON(400, gin.H{"error": "Bad parameters (wrong role type)"})
			return
		}
		if err := CreateUser(newUser); err != nil {
			c.JSON(400, gin.H{"error": "Creation failed"})
			return
		}
		c.JSON(200, "User created")
	})
}