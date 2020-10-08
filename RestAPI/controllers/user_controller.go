package controllers

import (
	"net/http"
	"personalProjects/RestAPI/domain/users"
	"personalProjects/RestAPI/services"
	"personalProjects/RestAPI/utils/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("User ID should be an Integer value")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userID)
	if getErr != nil {
		err := errors.NewBadRequestError("User ID should be an Integer value")
		c.JSON(err.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: Handle Error
	// 	fmt.Println("Wrong Input")
	// }
	// fmt.Println(string(bytes))
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	//TODO Handle Error
	// 	return
	// }
	// user, err := services.CreateUser(user)
	// if err != nil {
	// 	//TODO Handle Error
	// 	return
	// }
	//The above code can be replaced by
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	res, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusOK, res)

}

func UpdateUserInfo(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func DeleteUserInfo(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
