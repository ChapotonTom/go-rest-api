package role

import (
	"database/sql"
)

type SQLite struct {
	DB *sql.DB
}

func (s *SQLite) FindAll() []Role {
	roles := []Role{}
	rows, _ := s.DB.Query("SELECT * FROM Role")
	defer rows.Close()
	for rows.Next() {
		role := Role{}
		rows.Scan(&role.Id, &role.Userid, &role.Type)
		roles = append(roles, role)
	}
	return roles
}

func (s *SQLite) FindByUserId(userId int) []Role {
	roles := []Role{}
	rows, _ := s.DB.Query("SELECT * FROM Role WHERE userId = ?", userId)
	defer rows.Close()
	for rows.Next() {
		role := Role{}
		rows.Scan(&role.Id, &role.Userid, &role.Type)
		roles = append(roles, role)
	}
	return roles
}

func (s *SQLite) Add(role Role) error {
    stmt, _ := s.DB.Prepare("INSERT INTO Role(userId, type) values (?, ?)")
	_, err := stmt.Exec(role.Userid, role.Type)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLite) DeleteByUserID(userId int) error {
    stmt, _ := s.DB.Prepare("DELETE FROM Role WHERE userId = ?")
	_, err := stmt.Exec(userId)
	if err != nil {
		return err
	}
	return nil
}

func NewRoleTable(connexion *sql.DB) *SQLite {
	stmt, _ := connexion.Prepare(`
		CREATE TABLE IF NOT EXISTS "Role" (
			"id"		INTEGER,
			"userId"	INTEGER,
			"type"		TEXT,
			FOREIGN 	KEY("userId") REFERENCES User("id"),
			PRIMARY 	KEY("id" AUTOINCREMENT)
		);
	`)
	stmt.Exec()
	return &SQLite{
		DB: connexion,
	}
}