package user

import (
	"restapi/roles"
)

type formatedUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Roles []string `json:"roles"`
}

func filterRoles(roles []role.Role, userId int) []string {
	rolesFiltered := []string{}
	for _, role := range roles {
		if role.Userid == userId {
			rolesFiltered = append(rolesFiltered, role.Type)
		}
	}
	return rolesFiltered
}

func GetUsers() []formatedUser {
	users, _ := FindAll()
	roles, _ := role.FindAll()

	formattedUsers := []formatedUser{}

	for _, user := range users {
		userRoles := filterRoles(roles, user.Id)
		newUser := formatedUser{
			Id: user.Id,
			Name: user.Name,
			Roles: userRoles,
		};
		formattedUsers = append(formattedUsers, newUser)
	}
	return formattedUsers
}

func GetSingleUser(userId int) formatedUser {
	user, _ := FindById(userId)
	roles, _ := role.FindByUserId(userId)
	return formatedUser{
		Id: user.Id,
		Name: user.Name,
		Roles: filterRoles(roles, user.Id),
	};
}