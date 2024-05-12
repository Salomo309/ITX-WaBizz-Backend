package models

/*
Struct: User ->
Represent the user row that is stored in database.
*/
type User struct {
	Email       string
	IsActive    bool
	IsAdmin     bool
	DeviceToken string
}

/*
Struct: LoginRequest ->
Structure that is used as a request body of an endpoint.
Needed on CheckUserLogin function.
 */
type LoginRequest struct {
	Email       string	`json:"Email"`
	DeviceToken string 	`json:"DeviceToken"`
}

/*
Struct: LoginResponse ->
Structure that is used as a response body of an endpoint.
Needed on CheckUserLogin function.
 */
type LoginResponse struct {
	Message		string 	`json:"Message"`
	IsAdmin		bool 	`json:"IsAdmin"`
}

/*
Struct: UserInsertRequest ->
Structure that is used as a request body of an endpoint.const
Needed on InsertUser function.
*/
type UserInsertRequest struct {
	Email		string	`json:"Email"`
	IsActive	bool	`json:"IsActive"`
}
