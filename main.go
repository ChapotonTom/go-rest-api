package main

import (
    "restapi/config"
    "restapi/users"
    "restapi/roles"
    "github.com/gin-gonic/gin"
)

func main() {
    db, _ := config.GetDB()
    user.NewUsers(db)
    role.NewRoles(db)
    
    router := gin.Default()

    router.Use(gin.Logger())
	router.Use(gin.Recovery())

    user.UserRouter(router)

    router.Run("localhost:8080")
}
