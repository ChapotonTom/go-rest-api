package user

import (
	"database/sql"
)

type SQLite struct {
	DB *sql.DB
}

func (s *SQLite) GetAll() []User {
	users := []User{}
	rows, _ := s.DB.Query("SELECT * FROM User")
	defer rows.Close()
	for rows.Next() {
		u := User{}
		rows.Scan(&u.Id, &u.Name, &u.Password)
		users = append(users, u)
	}
	return users
}

func (s *SQLite) FindById(id int) (*User, error) {
	user := User{}

	err := s.DB.QueryRow("SELECT id, name, password FROM user WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SQLite) Add(user User) error {
    stmt, _ := s.DB.Prepare("INSERT INTO user(name, password) values (?, ?)")
	_, err := stmt.Exec(user.Name, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func NewUsers(connexion *sql.DB) *SQLite {
	stmt, _ := connexion.Prepare(`
		CREATE TABLE IF NOT EXISTS "user" (
			"id"	INTEGER UNIQUE,
			"name"	TEXT UNIQUE,
			"password"	TEXT,
			PRIMARY KEY("id" AUTOINCREMENT)
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: connexion,
	}
}