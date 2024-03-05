package models

type AccessTokenRequest struct {
	Google_ID string
}

type RefreshToken struct {
	Google_ID     string
	Refresh_Token string
}

type AccessTokenResponse struct {
	Google_ID     string	`json:"google_id"`
	Access_Token  string	`json:"access_token"`
}