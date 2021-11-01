package auth

import (
	"restapi/users"
	"restapi/roles"
	"restapi/utils"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func formatUserRoles(userRoles []role.Role) []string {
	roles := []string{}
	for _, role := range userRoles {
		roles = append(roles, role.Type)
	}
	return roles
}

func HandleLogin(c *gin.Context) {
	var login Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": "Bad parameters"})
		return
	}
	userInfos, err := user.FindByName(login.Username)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	roles, _ := role.FindByUserId(userInfos.Id)
	if !utils.CheckPasswordHash(login.Password, userInfos.Password) {
		c.JSON(401, gin.H{"error": "Wrong login informations"})
		return
	}
	token, _ := CreateToken(userInfos.Id, userInfos.Name, formatUserRoles(roles))
	c.JSON(200, gin.H{"token": token})
}