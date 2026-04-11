package dto

type AccountDTO struct {
	Username  string  `json:"username" binding:"required,min=3,max=30"`
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=8"`
	Language  *string `json:"language"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type ResetPasswordDTO struct {
	Identifier  string `json:"identifier" binding:"required"` // can be email or username
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}
