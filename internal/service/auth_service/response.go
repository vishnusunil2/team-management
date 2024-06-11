package auth_service

type UserResponse struct {
	Id          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
