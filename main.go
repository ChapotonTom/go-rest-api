package main

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./company.db")

    stmt, err := db.Prepare(`
        CREATE TABLE IF NOT EXISTS "user" (
            "id"	INTEGER,
            "name"	TEXT,
            "password"	TEXT,
            PRIMARY KEY("id" AUTOINCREMENT)
        );
    `)
    stmt.Exec()
}
