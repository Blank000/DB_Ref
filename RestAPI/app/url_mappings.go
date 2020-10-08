package app

import "personalProjects/RestAPI/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:user_id", controllers.GetUserInfo)
	router.DELETE("/users/:user_id", controllers.DeleteUserInfo)
	router.POST("/users/user_id", controllers.UpdateUserInfo)
}
