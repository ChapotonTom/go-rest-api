package role

import (
	"restapi/database"
)

func FindAll() ([]Role, error) {
	roles := []Role{}
	rows, err := database.DBCon.Query("SELECT * FROM Role")
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
	roles := []Role{}
	rows, err := database.DBCon.Query("SELECT * FROM Role WHERE userId = ?", userId)
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

func Add(userId int, role string) error {
    stmt, _ := database.DBCon.Prepare("INSERT INTO Role(userId, type) values (?, ?)")
	_, err := stmt.Exec(userId, role)
	if err != nil {
		return err
	}
	return nil
}

func Delete(roleId int) error {
    stmt, _ := database.DBCon.Prepare("DELETE FROM Role WHERE id = ?")
	_, err := stmt.Exec(roleId)
	if err != nil {
		return err
	}
	return nil
}

func NewRoles() {
	stmt, _ := database.DBCon.Prepare(`
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