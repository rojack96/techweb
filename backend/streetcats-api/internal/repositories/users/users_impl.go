package users

import (
	"context"
	"streetcats-api/internal/entities"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usersRepositoryImpl struct {
	pg  *pgxpool.Pool
	ctx context.Context
}

func (r *usersRepositoryImpl) SetContext(ctx any) {
	switch v := ctx.(type) {
	case *gin.Context:
		r.ctx = v.Request.Context()
	case context.Context:
		r.ctx = v
	default:
		r.ctx = context.Background()
	}
}

func NewUsersRepository(pg *pgxpool.Pool) Repository {
	return &usersRepositoryImpl{pg: pg}
}

func (r *usersRepositoryImpl) CreateUser(username, email string, language, firstName, lastName *string) (*entities.AccountProfile, error) {
	tx, err := r.pg.BeginTx(r.ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(r.ctx)
		} else {
			_ = tx.Commit(r.ctx)
		}
	}()

	// Insert account
	var (
		accountID            uint64
		createdAt, updatedAt time.Time
	)

	err = tx.QueryRow(r.ctx,
		`INSERT INTO users.accounts (username, email, language)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`,
		username, email, language).Scan(&accountID, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}

	// Insert profile if firstName or lastName is provided

	if firstName != nil || lastName != nil {
		var pid uint64

		err = tx.QueryRow(r.ctx,
			`INSERT INTO users.profiles (user_id, first_name, last_name)
			VALUES ($1, $2, $3)
			RETURNING id`, accountID, firstName, lastName).
			Scan(&pid)
		if err != nil {
			return nil, err
		}

	}

	account := entities.Account{
		ID:        accountID,
		Username:  username,
		Email:     email,
		Language:  language,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	profile := entities.Profile{
		UserID:    accountID,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	accountProfile := &entities.AccountProfile{Account: account, Profile: profile}

	return accountProfile, nil
}
