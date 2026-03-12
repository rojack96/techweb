package dto

type EmailInfoDTO struct {
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Email    string  `json:"email"`
	Username *string `json:"username"`
	Language *string `json:"language"`
	NotificationDTO
}
