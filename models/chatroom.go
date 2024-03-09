package models

type ChatList struct {
	CustomerName	string
	Timendate    	string
	IsRead			*bool
	StatusRead		*string
	Content			string
	MessageType		string
}