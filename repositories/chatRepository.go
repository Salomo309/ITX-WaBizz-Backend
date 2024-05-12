package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

/*** Chat Repository Declaration and Implementation***/

// Interface for chat repository
type ChatRepository interface {
	CreateChat(*models.Chat) error
	GetChats(int) ([]models.Chat, error)
	MarkAllChatsAsRead(int) error
}

// Implementation of chat repository
type chatRepo struct {
	db         		*sql.DB
	createStmt 		*sql.Stmt
	getStmt    		*sql.Stmt
	updateReadStmt	*sql.Stmt
}

// Function to create new chat repository. Prepare all statement and return the instance.
func NewChatRepository(db *sql.DB) (ChatRepository, error) {
	createStmt, err := db.Prepare("INSERT INTO Chat (chat_id, email, chatroom_id, timendate, isRead, statusRead, content, messageType) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	getStmt, err := db.Prepare("SELECT chat_id, email, chatroom_id, timendate, isRead, statusRead, content, messageType FROM Chat WHERE chatroom_id = ?")
	if err != nil {
		return nil, err
	}
	
	updateReadStmt, err := db.Prepare("UPDATE Chat SET isRead = '1' WHERE chatroom_id = ? AND isRead = '0'")
	if err != nil {
		return nil, err
	}

	return &chatRepo{
		db:         db,
		createStmt: createStmt,
		getStmt:    getStmt,
		updateReadStmt: updateReadStmt,
	}, nil
}

/*** Chat Repository Function Implementation ***/

// Insert new chat to database.
func (repo *chatRepo) CreateChat(chat *models.Chat) error {
	_, err := repo.createStmt.Exec(&chat.ChatID, &chat.Email, &chat.ChatroomID, &chat.Timendate, &chat.IsRead, &chat.StatusRead, &chat.Content, &chat.MessageType)
	if err != nil {
		return err
	}

	return nil
}

// Find all chats in a spesific chatroom.
func (repo *chatRepo) GetChats(chatroomID int) ([]models.Chat, error) {
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

// Update chats that has not been read into read.
func (repo *chatRepo) MarkAllChatsAsRead(chatroomID int) error {
	_, err := repo.updateReadStmt.Exec(chatroomID)
	if err != nil {
		return err
	}

	return nil
}
