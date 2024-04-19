package models

type Reply struct {
	ReplyId		 int	`json:"ReplyID"`
	UserId		 int	`json:"UserID"`
	ChatroomId	 int 	`json:"ChatroomID"`
	Timendate	 string `json:"Timendate"`
	content		 string	`json:"Content"`
	statusRead	 string	`json:"StatusRead"`
}