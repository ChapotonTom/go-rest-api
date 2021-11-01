package user

import (
	"errors"
	"restapi/roles"
	"restapi/utils"
)

type FormatedUser struct {
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

func GetUsers() []FormatedUser {
	users, _ := FindAll()
	roles, _ := role.FindAll()

	formattedUsers := []FormatedUser{}

	for _, user := range users {
		userRoles := filterRoles(roles, user.Id)
		newUser := FormatedUser{
			Id: user.Id,
			Name: user.Name,
			Roles: userRoles,
		};
		formattedUsers = append(formattedUsers, newUser)
	}
	return formattedUsers
}

func GetSingleUser(userId int) FormatedUser {
	user, _ := FindById(userId)
	roles, _ := role.FindByUserId(userId)
	return FormatedUser{
		Id: user.Id,
		Name: user.Name,
		Roles: filterRoles(roles, user.Id),
	};
}

func CreateUser(newUser User) error {
	hash, err := utils.HashPassword(newUser.Password)
	if err != nil {
		return errors.New("User creation failed")
	}
	newUser.Password = hash
	userId, err := Add(newUser)
	if err != nil {
		return errors.New("User creation failed")
	}
	for _, userRole := range newUser.Roles {
		if err := role.Add(userId, userRole); err != nil {
			return errors.New("User's role creation failed")
		}
	}
	return nil
}