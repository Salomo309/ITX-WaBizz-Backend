package models

type Reply struct {
	ReplyId		 int	`json:replyId`
	UserId		 int	`json:userId`
	ChatroomId	 int 	`json:chatroomId`
	Timendate	 string `json:timendate`
	content		 string	`json:content`
	statusRead	 string	`json:statusRead`
}