package models

/*
Struct: Chat ->
Represent the chat row that is stored in database.
*/
type Chat struct {
	ChatID		 int		`json:"ChatID"`
	Email		 *string	`json:"Email"`
	ChatroomID	 int 		`json:"ChatroomID"`
	Timendate	 string 	`json:"Timendate"`
	IsRead		 *string	`json:"IsRead"`
	StatusRead	 *string	`json:"StatusRead"`
	Content		 string		`json:"Content"`
	MessageType	 string		`json:"MessageType"`
}
