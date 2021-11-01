package role

type Role struct {
	Id int `json:"id"`
	Userid int `json:"userId"`
	Type string `json:"type"`
}

type RolesUpdate struct {
	Roles []string `json:"roles" binding:"max=2"`
}