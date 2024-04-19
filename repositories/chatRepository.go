package repositories

import (
	"database/sql"
	"fmt"
	"itx-wabizz/models"
)

type ChatRepository interface {
	CreateChat(chat *models.Chat) error
	GetChat(chatId int) (*models.Chat, error)
}

type MySQLChatRepository struct {
	db         *sql.DB
	createStmt *sql.Stmt
	getStmt    *sql.Stmt
}

func NewMySQLChatRepository(db *sql.DB) (*MySQLChatRepository, error) {
	// Prepare SQL statements
	createStmt, err := db.Prepare("INSERT INTO Chat (chatroomId, timendate, isRead, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing create statement:", err)
		return nil, err
	}

	getStmt, err := db.Prepare("SELECT chatId, chatroomId, timendate, isRead, content FROM Chat WHERE chatId = ?")
	if err != nil {
		fmt.Println("Error preparing get statement:", err)
		return nil, err
	}

	return &MySQLChatRepository{
		db:         db,
		createStmt: createStmt,
		getStmt:    getStmt,
	}, nil
}

func (repo *MySQLChatRepository) CreateChat(chat *models.Chat) error {
	_, err := repo.createStmt.Exec(chat.ChatroomId, chat.Timendate, chat.IsRead, chat.Content)
	if err != nil {
		fmt.Println("Error executing create statement:", err)
		return err
	}
	return nil
}

func (repo *MySQLChatRepository) GetChat(chatId int) (*models.Chat, error) {
	chat := &models.Chat{}
	err := repo.getStmt.QueryRow(chatId).Scan(&chat.ChatId, &chat.ChatroomId, &chat.Timendate, &chat.IsRead, &chat.Content)
	if err != nil {
		fmt.Println("Error retrieving chat:", err)
		return nil, err
	}
	return chat, nil
}
