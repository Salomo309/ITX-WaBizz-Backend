package repositories

import (
	"database/sql"

	"itx-wabizz/models"
)

/*** Chatroom Repository Declaration and Implementation***/

// Interface for chatroom repository
type ChatroomRepository interface {
	GetChatroomList() ([]models.ChatList, error)
	SearchChatroomByContact(string) ([]models.ChatList, error)
	SearchChatroomByMessage(string) ([]models.MessageSearchResult, error)
	Insert(*models.Chatroom) error
	GetChatroomByPhone(string) (*models.Chatroom, error)
	GetChatroomByID(int) (*models.Chatroom, error)
}

// Implementation of chatroom repository
type chatroomRepo struct {
	db                     	*sql.DB
	getChatListStmt        	*sql.Stmt
	searchByContactStmt    	*sql.Stmt
	searchByMessageStmt    	*sql.Stmt
	insertStmt             	*sql.Stmt
	getChatroomByPhoneStmt 	*sql.Stmt
	getChatroomByIDStmt    	*sql.Stmt
}

// Function to create new chatroom repository. Prepare all statement and return the instance.
func NewChatroomRepository(db *sql.DB) (ChatroomRepository, error) {
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
		return nil, err
	}

	searchByContactStmt, err := db.Prepare(`
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
		RowNum = 1 AND CustomerName LIKE CONCAT('%', ?, '%')`)
	if err != nil {
		return nil, err
	}

	searchByMessageStmt, err := db.Prepare(`
	SELECT CustomerName, Timendate, Content
	FROM (
		SELECT
			Chatroom.customer_name AS CustomerName,
			Chat.timendate AS Timendate,
			Chat.content AS Content,
			ROW_NUMBER() OVER (PARTITION BY Chatroom.chatroom_id ORDER BY Chat.timendate DESC) AS RowNum
		FROM
			Chatroom
		LEFT JOIN Chat ON Chatroom.chatroom_id = Chat.chatroom_id
		WHERE
			Chat.content LIKE CONCAT('%', ?, '%')
	) AS Subquery
	WHERE RowNum = 1`)
	if err != nil {
		return nil, err
	}

	insertStmt, err := db.Prepare("INSERT INTO Chatroom (customer_phone, customer_name) values (?, ?)")
	if err != nil {
		return nil, err
	}

	getChatroomByPhoneStmt, err := db.Prepare("SELECT chatroom_id, customer_phone, customer_name FROM Chatroom WHERE customer_phone = ?")
	if err != nil {
		return nil, err
	}

	getChatroomByIDStmt, err := db.Prepare("SELECT chatroom_id, customer_phone, customer_name FROM Chatroom WHERE chatroom_id = ?")
	if err != nil {
		return nil, err
	}

	return &chatroomRepo{
		db:                     db,
		getChatListStmt:        getChatListStmt,
		searchByContactStmt:    searchByContactStmt,
		searchByMessageStmt:    searchByMessageStmt,
		insertStmt:             insertStmt,
		getChatroomByPhoneStmt: getChatroomByPhoneStmt,
		getChatroomByIDStmt:    getChatroomByIDStmt,
	}, nil
}

/*** Chatroom Repository Function Implementation ***/

// Get a list of chatroom data to be viewed by user
func (repo *chatroomRepo) GetChatroomList() ([]models.ChatList, error) {
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
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatlists, nil
}

// Find all chatroom where the customer name has a substring that match.
func (repo *chatroomRepo) SearchChatroomByContact(searchText string) ([]models.ChatList, error) {
	rows, err := repo.searchByContactStmt.Query(searchText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatlists []models.ChatList

	for rows.Next() {
		var chatlist models.ChatList
		err := rows.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.IsRead, &chatlist.StatusRead, &chatlist.Content, &chatlist.MessageType, &chatlist.CountUnread)
		if err != nil {
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatlists, nil
}

// Find all chatroom where a message content has a substring that match.
func (repo *chatroomRepo) SearchChatroomByMessage(searchText string) ([]models.MessageSearchResult, error) {
	rows, err := repo.searchByMessageStmt.Query(searchText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chatlists []models.MessageSearchResult

	for rows.Next() {
		var chatlist models.MessageSearchResult
		err := rows.Scan(&chatlist.CustomerName, &chatlist.Timendate, &chatlist.Content)
		if err != nil {
			return nil, err
		}
		chatlists = append(chatlists, chatlist)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chatlists, nil
}

// Insert new chatroom to database.
func (repo *chatroomRepo) Insert(chatroom *models.Chatroom) error {
	_, err := repo.insertStmt.Exec(chatroom.CustomerPhone, chatroom.CustomerName)
	if err != nil {
		return err
	}

	return nil
}

// Find a chatroom from database using spesific customer phone.
func (repo *chatroomRepo) GetChatroomByPhone(customerPhone string) (*models.Chatroom, error) {
	row := repo.getChatroomByPhoneStmt.QueryRow(customerPhone)

	var chatroom models.Chatroom
	err := row.Scan(&chatroom.ChatroomID, &chatroom.CustomerPhone, &chatroom.CustomerName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &chatroom, nil
}

// Find a chatroom from database using spesific chatroom ID.
func (repo *chatroomRepo) GetChatroomByID(chatroomID int) (*models.Chatroom, error) {
	row := repo.getChatroomByIDStmt.QueryRow(chatroomID)

	var chatroom models.Chatroom
	err := row.Scan(&chatroom.ChatroomID, &chatroom.CustomerPhone, &chatroom.CustomerName)
	if err != nil {
		return nil, err
	}

	return &chatroom, nil
}
