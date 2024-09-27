package service

import (
	"errors"
	"fmt"
	"task-management/database"
	"task-management/model/user"
)

type UserService struct {}

func (u *UserService) SignUp(user user.User)(bool,error){
	queryErr := database.DB.Table("users").Create(&user)
	if queryErr.Error!=nil {
		fmt.Println("queryErr",queryErr.Error)
		return false,queryErr.Error
	}
	return true,nil
}
func (u *UserService) SignIn(userInfo user.UserSignIn) (user.User, error) {

	var fetchedUser user.User
	queryErr := database.DB.Table("users").Where("username = ?", userInfo.UserName).Scan(&fetchedUser)
	
	if queryErr.Error != nil {
		fmt.Println("queryErr", queryErr.Error)
		return fetchedUser, queryErr.Error
	}

	
	
	if fetchedUser.UserName != "" && fetchedUser.Password == userInfo.Password {
		return fetchedUser, nil
	}

	return fetchedUser, errors.New("user not found")
}
