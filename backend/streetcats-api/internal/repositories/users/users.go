package users

import "streetcats-api/internal/entities"

type Repository interface {
	CreateUser(username, email string, language, firstName, lastName *string) (*entities.AccountProfile, error)
	//GetUserByUsername(username string) (*entities.AccountProfile, error)
}
