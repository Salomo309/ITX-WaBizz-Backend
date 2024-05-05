package repositories

import (
	"database/sql"
	"fmt"
	"itx-wabizz/models"
)

type ChatRepository interface {
	CreateChat(chat *models.Chat) error
	GetChats(chatroomID int) (*models.Chat, error)
}

type MySQLChatRepository struct {
	db         *sql.DB
	createStmt *sql.Stmt
	getStmt    *sql.Stmt
}

func NewMySQLChatRepository(db *sql.DB) (*MySQLChatRepository, error) {
	// Prepare SQL statements
	createStmt, err := db.Prepare("INSERT INTO Chat (chat_id, email, chatroom_id, timendate, isRead, statusRead, content, messageType) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing create statement:", err)
		return nil, err
	}

	getStmt, err := db.Prepare("SELECT chat_id, email, chatroom_id, timendate, isRead, statusRead, content, messageType FROM Chat WHERE chatroom_id = ?")
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
	_, err := repo.createStmt.Exec(&chat.ChatID, &chat.Email, &chat.ChatroomID, &chat.Timendate, &chat.IsRead, &chat.StatusRead, &chat.Content, &chat.MessageType)
	if err != nil {
		fmt.Println("Error executing create statement:", err)
		return err
	}
	return nil
}

func (repo *MySQLChatRepository) GetChats(chatroomID int) ([]models.Chat, error) {
	rows, err := repo.getStmt.Query(chatroomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []models.Chat

	for rows.Next() {
		var chat models.Chat
		err := rows.Scan(&chat.ChatID, &chat.Email, &chat.ChatroomID, &chat.Timendate, &chat.IsRead, &chat.StatusRead, &chat.Content, &chat.MessageType)
		if err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

func (repo *MySQLChatRepository) MarkAllChatsAsRead(chatroomID int) error {
	_, err := repo.db.Exec("UPDATE Chat SET isRead = ? WHERE chatroom_id = ? AND isRead = ?", "1", chatroomID, "0")
	if err != nil {
		fmt.Println("Error executing update statement:", err)
		return err
	}
	return nil
}
