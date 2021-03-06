package user

type User struct {
	Id int `json:"id"`
	Name string `json:"name" binding:"required,max=50"`
	Password string `json:"password" binding:"required,max=50"`
	Roles []string `json:"roles" binding:"max=2"`
}

type FormatedUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Roles []string `json:"roles"`
}