package users

import (
	"fmt"
	"personalProjects/RestAPI/datasources/mysql/userdb"
	"personalProjects/RestAPI/utils/dateutils"
	"personalProjects/RestAPI/utils/errors"
	"strings"
)

const (
	queryToInsertNewUser   = "INSERT INTO users_profile(email, password, first_name, second_name, sex, sexual_orientation, date_created, age, mobile) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? );"
	querytToQueryAboutUser = "SELECT * FROM users_profile"
)

var (
	indexUniqueEmail = "email_UNIQUE"
)

func (user *User) Save() *errors.RestErr {

	user.DateCreated = dateutils.GetNowString()
	fmt.Println("date Created" + user.DateCreated)
	stmt, err := userdb.Client.Prepare(queryToInsertNewUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	fmt.Println(user.Email, user.Password, user.FirstName, user.LasttName, user.Sex, user.SexualOrientation, user.DateCreated, user.Age, user.MobileNumber)
	insertRes, err := stmt.Exec(user.Email, user.Password, user.FirstName, user.LasttName, user.Sex, user.SexualOrientation, user.DateCreated, user.Age, user.MobileNumber)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when tryin to save user: %s", err.Error()))
	}
	userID, err := insertRes.LastInsertId()
	// if insertRes != nil {
	// 	if insertRes.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s is already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists in the system", user.UserID))
	// }
	user.UserID = userID

	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := userdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := userdb.Client.Prepare(queryToInsertNewUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.UserID))
	}
	user.UserID = result.UserID
	user.FirstName = result.FirstName
	user.LasttName = result.LasttName
	user.Age = result.Age
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	user.MobileNumber = result.MobileNumber
	user.Password = result.Password
	user.Sex = result.Sex
	user.SexualOrientation = result.SexualOrientation
	return nil
}
