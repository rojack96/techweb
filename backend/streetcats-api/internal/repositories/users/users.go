package users

import (
	"context"
	"streetcats-api/internal/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, account entities.Account, profile entities.Profile) (*entities.AccountProfile, error)
	//GetUserByUsername(username string) (*entities.AccountProfile, error)
}
