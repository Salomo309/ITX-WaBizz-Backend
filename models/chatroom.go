package models

/*
Struct: Chatroom ->
Represent the chatroom that is stored in database.
*/
type Chatroom struct {
    ChatroomID      int     `json:"ChatroomID"`
    CustomerPhone   string  `json:"CustomerPhone"`
    CustomerName    string  `json:"CustomerName"`
}

/*
Struct: Chatlist ->
Represent the element of chatroom list that will be viewed by user.
*/
type ChatList struct {
	CustomerName    string     `json:"customerName"`
    Timendate       string     `json:"timendate"`
    IsRead          *string    `json:"isRead"`
    StatusRead      *string    `json:"statusRead"`
    Content         string     `json:"content"`
	MessageType	    string     `json:"messageType"`
    CountUnread     int        `json:"countUnread"`
}

/*
Struct: MessageSearchResult ->
Needed to represent result of chatroom list searching by message.
*/
type MessageSearchResult struct {
	CustomerName    string     `json:"customerName"`
    Timendate       string     `json:"timendate"`
    Content         string     `json:"content"`
}