package dto

type AccountDTO struct {
	Username  string  `json:"username" binding:"required,min=3,max=30"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=8"`
	Language  *string `json:"language"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type ResetPasswordDTO struct {
	Identifier  string `json:"identifier" binding:"required"` // can be email or username
	NewPassword string `json:"new_password" binding:"required,min=8"`
}
