package middleware

import (
	"fmt"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

func checkType(roleType string, roles []string) bool {
	for _, role := range roles {
		if role == roleType {
			return true
		}
	}
	return false
}

func convertRoleClaims(tokenClaims jwt.Claims) []string {
	aInterface := tokenClaims.(jwt.MapClaims)["roles"].([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
    	aString[i] = v.(string)
	}
	return aString
}

func validatetoken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		  return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	   }
	   return []byte("SECRET_KEY"), nil // normally should get variable in env
	})
	if err != nil {
		fmt.Println(err)
	   return nil, err
	}
	return token, nil
}

func authorizeJWT(c *gin.Context) (jwt.Claims, error) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("Unauthorized")
	}
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, err := validatetoken(tokenString)
	if err != nil || !token.Valid {
		return nil, errors.New("Unauthorized")
	}
	tokenClaims := token.Claims
	return tokenClaims, nil
}

func AuthorizeUser(c *gin.Context) {
	tokenClaims, err := authorizeJWT(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(401)
	}
	c.Set("userId", tokenClaims.(jwt.MapClaims)["userId"])
}

func AuthorizeManager(c *gin.Context) {
	tokenClaims, err := authorizeJWT(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(401)
	}
	userRoles := convertRoleClaims(tokenClaims)
	isManager := checkType("manager", userRoles)
	if !isManager {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(401)
	}
}

func AuthorizeEmployee(c *gin.Context) {
	tokenClaims, err := authorizeJWT(c)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(401)
	}
	userRoles := convertRoleClaims(tokenClaims)
	isEmployee := checkType("employee", userRoles)
	if !isEmployee {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.AbortWithStatus(401)
	}
}