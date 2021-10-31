package role

import (
	"restapi/config"
	"database/sql"
)

func FindAll() ([]Role, error) {
	db, _ := config.GetDB()
	roles := []Role{}
	rows, err := db.Query("SELECT * FROM Role")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		role := Role{}
		rows.Scan(&role.Id, &role.Userid, &role.Type)
		roles = append(roles, role)
	}
	return roles, nil
}

func FindByUserId(userId int) ([]Role, error) {
	db, _ := config.GetDB()
	roles := []Role{}
	rows, err := db.Query("SELECT * FROM Role WHERE userId = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		role := Role{}
		rows.Scan(&role.Id, &role.Userid, &role.Type)
		roles = append(roles, role)
	}
	return roles, nil
}

func Add(role Role) error {
	db, _ := config.GetDB()
    stmt, _ := db.Prepare("INSERT INTO Role(userId, type) values (?, ?)")
	_, err := stmt.Exec(role.Userid, role.Type)
	if err != nil {
		return err
	}
	return nil
}

func DeleteByUserID(userId int) error {
	db, _ := config.GetDB()
    stmt, _ := db.Prepare("DELETE FROM Role WHERE userId = ?")
	_, err := stmt.Exec(userId)
	if err != nil {
		return err
	}
	return nil
}

func NewRoles(connexion *sql.DB) {
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
}