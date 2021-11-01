package user

import (
	"restapi/database"
)

func FindAll() ([]User, error) {
	users := []User{}
	rows, err := database.DBCon.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}
	return users, nil
}

func FindAllExceptOne(userId int) ([]User, error) {
	users := []User{}
	rows, err := database.DBCon.Query("SELECT * FROM User WHERE id != ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Name, &user.Password)
		users = append(users, user)
	}
	return users, nil
}

func FindById(id int) (*User, error) {
	user := User{}
	err := database.DBCon.QueryRow("SELECT * FROM User WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByName(username string) (*User, error) {
	user := User{}
	err := database.DBCon.QueryRow("SELECT * FROM User WHERE name = ?", username).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Add(user User) (int, error) {
    stmt, _ := database.DBCon.Prepare("INSERT INTO User(name, password) values (?, ?)")
	result, err := stmt.Exec(user.Name, user.Password)
	if err != nil {
		return -1, err
	}
	var id int64
	id, err = result.LastInsertId()
	return int(id), nil
}

func NewUsers() {
		stmt, _ := database.DBCon.Prepare(`
		CREATE TABLE IF NOT EXISTS "User" (
			"id"	INTEGER UNIQUE,
			"name"	TEXT UNIQUE,
			"password"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	stmt.Exec()
}