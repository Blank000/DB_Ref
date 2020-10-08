package main

import (
	"personalProjects/RestAPI/app"
	_ "personalProjects/mysql"
)

func main() {
	app.Start()
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	// 	"muffy",
	// 	"pratapalokraj",
	// 	"127.0.0.1",
	// 	"users_db",
	// )
	// Client, err := sql.Open("mysql", dataSourceName)
	// if err != nil {
	// 	panic(err)
	// }
	// if err = Client.Ping(); err != nil {
	// 	panic(err)
	// }
	// log.Println("database successfully configured")
}
