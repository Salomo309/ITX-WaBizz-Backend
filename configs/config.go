package configs

import (
	"os"
)

var (
	// Variable to connect to database
	MysqlHost     string
	MysqlUser     string
	MysqlPassword string
	MysqlDatabase string

	// Additional variable
	StoragePath              string
	MessagingCredentialsPath string
	InfobipReceiveEndpoint   string
	ServerReceiveEndpoint    string
)

func InitConfiguration() {
	// Variable to connect to database
	MysqlHost = os.Getenv("MYSQL_HOST")
	MysqlUser = os.Getenv("MYSQL_USER")
	MysqlPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	MysqlDatabase = os.Getenv("MYSQL_DATABASE")

	// Additional variable
	StoragePath = os.Getenv("STORAGE_PATH")
	MessagingCredentialsPath = os.Getenv("MESSAGING_CREDENTIALS_PATH")
	InfobipReceiveEndpoint = os.Getenv("INFOBIP_RECEIVE_ENDPOINT")
	ServerReceiveEndpoint = os.Getenv("SERVER_RECEIVE_ENDPOINT")
}
