package repositories

import (
	"database/sql"
	"fmt"
	"itx-wabizz/models"
)

type ChatListRepository interface {
	GetChatList() ([]models.ChatList, error)
}

type MySQLChatListRepository struct {
	db				*sql.DB
	getChatListStmt	*sql.Stmt
}

func NewMySQLChatListRepository(db *sql.DB) (*MySQLChatListRepository, error){
	getChatListStmt, err := db.Prepare(`SELECT
    CustomerName,
    Timendate,
    IsRead,
    StatusRead,
    Content,
    MessageType,
    CountUnread
FROM (
    SELECT
        Chatroom.customer_name AS CustomerName,
        Chat.timendate AS Timendate,
        Chat.isRead AS IsRead,
        Chat.statusRead AS StatusRead,
        Chat.content AS Content,
        Chat.messageType AS MessageType,
        COALESCE(res2.countUnread, 0) AS CountUnread,
        ROW_NUMBER() OVER (PARTITION BY Chatroom.chatroom_id ORDER BY Chat.timendate DESC) AS RowNum
    FROM
        Chat
    JOIN
        Chatroom ON Chat.chatroom_id = Chatroom.chatroom_id
    LEFT JOIN
        (SELECT chatroom_id, COUNT(chatroom_id) AS countUnread FROM Chat WHERE isRead = "0" GROUP BY chatroom_id) AS res2
    ON
        Chat.chatroom_id = res2.chatroom_id
) AS Subquery
WHERE
    RowNum = 1;`)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}

	return &MySQLChatListRepository{
		db:				db,
		getChatListStmt:getChatListStmt,
	}, nil
}

func (repo *MySQLChatListRepository) GetChatList() ([]models.ChatList, error){
	rows, err := repo.getChatListStmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatlists []models.ChatList

	for rows.Next() {
		var chatlist models.ChatList
		err := rows.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.IsRead, &chatlist.StatusRead, &chatlist.Content, &chatlist.MessageType, &chatlist.CountUnread)
		if err != nil {
			fmt.Println("Error rows:", err)
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error rows2:", err)
		return nil, err
	}

	return chatlists, nil
}

func (repo *MySQLChatListRepository) SearchChatListByContact() (searchText string) ([]models.ChatList, error){
	query := `
		SELECT CustomerName, Timendate,	IsRead,	StatusRead, Content, MessageType, CountUnread
		FROM (
			SELECT
				Chatroom.customer_name AS CustomerName,
				Chat.timendate AS Timendate,
				Chat.isRead AS IsRead,
				Chat.statusRead AS StatusRead,
				Chat.content AS Content,
				Chat.messageType AS MessageType,
				COALESCE(res2.countUnread, 0) AS CountUnread,
				ROW_NUMBER() OVER (PARTITION BY Chatroom.chatroom_id ORDER BY Chat.timendate DESC) AS RowNum
			FROM
				Chat
			JOIN
				Chatroom ON Chat.chatroom_id = Chatroom.chatroom_id
			LEFT JOIN
				(SELECT chatroom_id, COUNT(chatroom_id) AS countUnread FROM Chat WHERE isRead = "0" GROUP BY chatroom_id) AS res2
			ON
				Chat.chatroom_id = res2.chatroom_id
		) AS Subquery
		WHERE
			RowNum = 1 AND CustomerName LIKE CONCAT('%', ?, '%')
    `
    rows, err := repo.db.Query(query, "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	var chatlists []models.ChatList

	for rows.Next() {
		var chatlist models.ChatList
		err := rows.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.IsRead, &chatlist.StatusRead, &chatlist.Content, &chatlist.MessageType, &chatlist.CountUnread)
		if err != nil {
			fmt.Println("Error rows:", err)
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error rows2:", err)
		return nil, err
	}

	return chatlists, nil
}

func (repo *MySQLChatListRepository) SearchChatListByMessage() (searchText string) ([]models.ChatList, error){
	query := `
		SELECT CustomerName, Timendate, Content
		FROM (
			SELECT
				Chatroom.customer_name AS CustomerName,
				COALESCE(Chat.timendate, Reply.timendate) AS Timendate,
				COALESCE(Chat.content, Reply.content) AS Content,
				ROW_NUMBER() OVER (PARTITION BY Chatroom.chatroom_id ORDER BY COALESCE(Chat.timendate, Reply.timendate) DESC) AS RowNum
			FROM
				Chatroom
			LEFT JOIN Chat ON Chatroom.chatroom_id = Chat.chatroom_id
			LEFT JOIN Reply ON Chatroom.chatroom_id = Reply.chatroom_id
			WHERE
				Chat.content LIKE CONCAT('%', ?, '%') OR
				Reply.content LIKE CONCAT('%', ?, '%')
		) AS Subquery
		WHERE RowNum = 1
    `
    rows, err := repo.db.Query(query, "%"+searchText+"%", "%"+searchText+"%", "%"+searchText+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

	var chatlists []models.ChatList

	for rows.Next() {
		var chatlist models.ChatList
		err := rows.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.IsRead, &chatlist.StatusRead, &chatlist.Content, &chatlist.MessageType, &chatlist.CountUnread)
		if err != nil {
			fmt.Println("Error rows:", err)
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error rows2:", err)
		return nil, err
	}

	return chatlists, nil
}