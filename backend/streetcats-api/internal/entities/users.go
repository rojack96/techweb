package entities

import "time"

type Account struct {
	ID        uint64    `db:"id"`
	Username  string    `db:"username"`
	Language  *string   `db:"language"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Profile struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	FirstName *string   `db:"first_name"`
	LastName  *string   `db:"last_name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AccountProfile struct {
	Account
	Profile
}
