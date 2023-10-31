package model

type UserCredentials struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserLogout struct {
	CustEmail string `json:"email"`
}
