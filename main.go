package main

import (
    "database/sql"
    "fmt"
    "restapi/users"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, _ := sql.Open("sqlite3", "./company.db")
    users := user.NewUsers(db)

    allUsers := users.GetAll()
    fmt.Println(allUsers)
}
