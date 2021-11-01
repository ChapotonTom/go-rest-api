package user

import (
	"strconv"
	"errors"
	"restapi/roles"
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

func HandleGetUsers(c *gin.Context) {
	userId, _ := c.Get("userId")
	id, _ := strconv.Atoi(userId.(string))
	users, err := GetOtherUsers(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Request Failed"})
		return
	}
	c.JSON(200, users)
}

func HandleGetUser(c *gin.Context) {
	userId, _ := c.Get("userId")
	id, _ := strconv.Atoi(userId.(string))
	user, err := GetSingleUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Request Failed"})
		return
	}
	c.JSON(200, user)
}

func HandleUserCreate(c *gin.Context) {
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
	c.JSON(200, gin.H{"success": "User created"})
}

func HandleUserUpdate(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Bad parameters"})
		return
	}
	var rolesUpdate role.RolesUpdate
	if err := c.ShouldBindJSON(&rolesUpdate); err != nil {
		c.JSON(400, gin.H{"error": "Bad parameters"})
		return
	}
	if err := validateRoles(rolesUpdate.Roles); err != nil {
		c.JSON(400, gin.H{"error": "Bad parameters (wrong role type)"})
		return
	}
	if _, err := FindById(userId); err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if err := UpdateUserRoles(userId, rolesUpdate.Roles); err != nil {
		c.JSON(500, gin.H{"error": "Update failed"})
		return
	}
	c.JSON(200, gin.H{"success": "User updated"})
}