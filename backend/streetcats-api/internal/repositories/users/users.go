package users

import "streetcats-api/internal/entities"

type Repository interface {
	CreateUser(username string, language, firstName, lastName *string) (*entities.AccountProfile, error)
	//GetUserByUsername(username string) (*entities.AccountProfile, error)
}
