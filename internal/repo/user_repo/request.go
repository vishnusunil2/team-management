package user_repo

type CreateUserRequest struct {
	Email     string
	Phone     string
	FirstName string
	LastName  string
	Password  string
}
