package users

import (
	"context"
	"streetcats-api/internal/entities"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usersRepositoryImpl struct {
	pg  *pgxpool.Pool
	ctx context.Context
}

func NewUsersRepository(pg *pgxpool.Pool) Repository {
	return &usersRepositoryImpl{pg: pg}
}

func (r *usersRepositoryImpl) CreateUser(ctx context.Context, account entities.Account, profile entities.Profile) (*entities.AccountProfile, error) {
	tx, err := r.pg.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	// Insert account
	err = tx.QueryRow(ctx,
		`INSERT INTO users.accounts (username, email, language)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`,
		account.Username, account.Email, account.Language).
		Scan(&account.ID, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Insert profile if firstName or lastName is provided
	if profile.FirstName != nil || profile.LastName != nil {
		var pid uint64

		err = tx.QueryRow(ctx,
			`INSERT INTO users.profiles (user_id, first_name, last_name)
			VALUES ($1, $2, $3)
			RETURNING id`,
			account.ID, profile.FirstName, profile.LastName).
			Scan(&pid)
		if err != nil {
			return nil, err
		}

	}

	accountProfile := &entities.AccountProfile{Account: account, Profile: profile}

	return accountProfile, nil
}
