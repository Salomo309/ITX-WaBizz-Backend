package models

type Chat struct {
	ChatId		 int	`json:chatId`
	ChatroomId	 int 	`json:chatroomId`
	Timendate	 string `json:timendate`
	isRead		 string	`json:isRead`
	content		 string	`json:content`
}