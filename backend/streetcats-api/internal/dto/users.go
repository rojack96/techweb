package dto

type AccountDTO struct {
	Username  string  `db:"username"`
	Language  *string `db:"language"`
	FirstName *string `db:"first_name"`
	LastName  *string `db:"last_name"`
}
