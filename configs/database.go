package configs

import (
	"os"
)

var (
	// Variable to connect to database
	MysqlHost = os.Getenv("MYSQL_HOST")
	MysqlUser = os.Getenv("MYSQL_USER")
	MysqlPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	MysqlDatabase = os.Getenv("MYSQL_DATABASE")
)
