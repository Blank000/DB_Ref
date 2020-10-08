package services

import (
	"fmt"
	"personalProjects/RestAPI/domain/users"
	"personalProjects/RestAPI/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := users.User{
		UserID: userID,
	}
	if getErr := result.Get(); getErr != nil {
		return nil, getErr
	}
	fmt.Println("THIS IS RESULT ")
	fmt.Println(result)
	return &result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.ValidateUser(); err != nil {
		return nil, err
	}
	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}
