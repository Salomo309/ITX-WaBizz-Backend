package models

type GoogleUserInfo struct {
	ID          string 	`json:"id"`
    Email       string 	`json:"email"`
    Name        string 	`json:"name"`
    Picture     string 	`json:"picture"`
}

type User struct {
	User_ID		int	
	Google_ID   string 
    Email       string 
    Name        string 
    Picture     string 
	Admin		bool
}

type UserResponseToken struct {
	Token		string 	`json:"token"`
	Email       string 	`json:"email"`
    Name        string 	`json:"name"`
    Picture     string 	`json:"picture"`
	Admin		bool	`json:"admin"`
}
