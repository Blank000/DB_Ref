package users

import (
	"personalProjects/RestAPI/utils/errors"
	"strings"
)

type User struct {
	UserID            int64  `json:"id"`
	Password          string `json:"password"`
	Email             string `json:"email"`
	MobileNumber      int64  `json:"mobile"`
	FirstName         string `json:"first_name"`
	LasttName         string `json:"last_name"`
	Age               int64  `json:"age"`
	Sex               string `json:"sex"`
	SexualOrientation string `json:"sexual_orientation"`
	DateCreated       string `json:"date_created"`
}

func (user *User) ValidateUser() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email Address")
	}
	return nil
}
