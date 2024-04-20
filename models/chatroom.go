package models

type ChatList struct {
	CustomerName string     `json:"customerName"`
    Timendate    string     `json:"timendate"`
    IsRead       *string    `json:"isRead"`
    StatusRead   *string    `json:"statusRead"`
    Content      string     `json:"content"`
	MessageType	 string     `json:"messageType"`
    CountUnread  int        `json:"countUnread"`
}

type MessageSearchResult struct {
	CustomerName string     `json:"customerName"`
    Timendate    string     `json:"timendate"`
    Content      string     `json:"content"`
}

type Chatroom struct {
    ChatroomID      int     `json:"ChatroomID"`
    CustomerPhone   string  `json:"CustomerPhone"`
    CustomerName    string  `json:"CustomerName"`
}