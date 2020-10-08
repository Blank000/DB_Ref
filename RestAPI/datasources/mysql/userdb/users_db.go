package userdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "personalProjects/mysql"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHostname = "mysql_users_hostname"
	mysqlUsersSchema   = "mysql_users_schema"
)

var (
	Client   *sql.DB
	userName = os.Getenv(mysqlUsersUsername)
	password = os.Getenv(mysqlUsersPassword)
	hostName = os.Getenv(mysqlUsersHostname)
	schema   = os.Getenv(mysqlUsersSchema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "muffy", "pratapalokraj", "127.0.0.1:3306", "users_db")
	fmt.Println("DATA" + dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	//Client = Client1
	fmt.Println("Client Confiigured")
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
