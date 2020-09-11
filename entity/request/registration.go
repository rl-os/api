package request

// CreateUser contain incoming data with user credentials
type CreateUser struct {
	Username string `json:"username" form:"user[username]" validate:"required"`
	Email    string `json:"email" form:"user[user_email]" validate:"required,email"`
	Password string `json:"password" form:"user[password]" validate:"required"`
}
