package models

type User struct {
    Email       string 
	IsActive	bool
}

type LoginToken struct {
	Email	string 	`json:"Email"`
}