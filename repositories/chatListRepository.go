package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

type chatListRepository interface {
	Insert(user *models.ChatList) error
}

type MySQLChatListRepository struct {
	db				*sql.DB
	getChatListStmt	*sql.Stmt
}

func NewMySQLChatListRepository(db *sql.DB) (*MySQLChatListRepository, error){
	getChatListStmt, err := db.Prepare(`SELECT Chatroom.customer_name AS CustomerName, res.timendate AS Timendate, Chat.isRead AS IsRead, Chat.statusRead AS StatusRead, Chat.content AS Content, Chat.messageType AS MessageType
	FROM
		(SELECT chatroom_id, MAX(timendate) AS timendate
		FROM Chat GROUP BY chatroom_id)
		AS res
	JOIN Chat ON res.timendate = Chat.timendate AND res.chatroom_id = Chat.chatroom_id
	JOIN Chatroom ON Chat.chatroom_id = Chatroom.chatroom_id;`)
	if err != nil {
		return nil, err
	}

	return &MySQLChatListRepository{
		db:				db,
		getChatListStmt:getChatListStmt,
	}, nil
}

func (repo *MySQLChatListRepository) GetChatList() (*models.ChatList, error){
	row := repo.getChatListStmt.QueryRow()

	var chatlist models.ChatList
	err := row.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.IsRead, &chatlist.StatusRead, &chatlist.Content, &chatlist.MessageType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &chatlist, nil
}