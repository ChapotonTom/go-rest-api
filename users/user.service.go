package user

import (
	"errors"
	"restapi/roles"
	"restapi/utils"
)

func filterRoles(roles []role.Role, userId int) []string {
	rolesFiltered := []string{}
	for _, role := range roles {
		if role.Userid == userId {
			rolesFiltered = append(rolesFiltered, role.Type)
		}
	}
	return rolesFiltered
}

func sortOutRoles(currentRoles []role.Role, newRoles []string) ([]string, []int) {
	rolesToCreate := []string{}
	rolesToDeleteIds := []int{}
	currentRolesString := []string{}

	for _, currentRole := range currentRoles {
		if !utils.StringInSlice(currentRole.Type, newRoles) {
			rolesToDeleteIds = append(rolesToDeleteIds, currentRole.Id)
		}
		currentRolesString = append(currentRolesString, currentRole.Type)
	}
	for _, newRole := range newRoles {
		if !utils.StringInSlice(newRole, currentRolesString) {
			rolesToCreate = append(rolesToCreate, newRole)
		}
	}
	return rolesToCreate, rolesToDeleteIds
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

func UpdateUserRoles(userId int, newUserRoles []string) error {
	userCurrentRoles, _ := role.FindByUserId(userId)
	rolesToCreate, rolesToDeleteIds := sortOutRoles(userCurrentRoles, newUserRoles)
	for _, roleToCreate := range rolesToCreate {
		if err := role.Add(userId, roleToCreate); err != nil {
			return errors.New("User's role update failed")
		}
	}

	for _, roleToDeleteId := range rolesToDeleteIds {
		if err := role.Delete(roleToDeleteId); err != nil {
			return errors.New("User's role update failed")
		}
	}
	return nil
}