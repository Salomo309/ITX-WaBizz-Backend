package repositories

import (
	"database/sql"
	"log"

	"itx-wabizz/configs"
)

var (
	// Database connection and important repositories that are needed throught app lifetime
	Db           *sql.DB
	UserRepo     UserRepository
	ChatlistRepo ChatroomRepository
	ChatRepo     ChatRepository
)

func InitDatabaseConnection() {
	var err error
	Db, err = sql.Open("mysql", configs.MysqlUser+":"+configs.MysqlPassword+"@tcp("+configs.MysqlHost+":3306)/"+configs.MysqlDatabase)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Ping database to check if the connection is valid
	if err := Db.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
}

func InitRepositories() {
	var err error
	UserRepo, err = NewUserRepository(Db)
	if err != nil {
		log.Fatal("Error initiating repositories: " + err.Error())
	}

	ChatlistRepo, err = NewChatroomRepository(Db)
	if err != nil {
		log.Fatal("Error initiating repositories: " + err.Error())
	}

	ChatRepo, err = NewChatRepository(Db)
	if err != nil {
		log.Fatal("Error initiating repositories: ")
	}
}
