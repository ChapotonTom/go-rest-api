package main

import (
    "restapi/database"
    "restapi/auth"
    "restapi/users"
    "restapi/roles"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "github.com/gin-gonic/gin"
)

func main() {
    var err error
    database.DBCon, err = sql.Open("sqlite3", "./company.db")
    if err != nil {
        dberr := fmt.Errorf("Database connection failed")
        fmt.Println(dberr.Error())
        return
    }
    user.NewUsers()
    role.NewRoles()
    
    router := gin.Default()

    router.Use(gin.Logger())
	router.Use(gin.Recovery())

    auth.AuthRouter(router)
    user.UserRouter(router)

    router.Run("localhost:8080")
}
