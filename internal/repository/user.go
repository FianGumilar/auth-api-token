package repository

import (
	"context"
	"database/sql"

	"github.com/FianGumilar/auth-api-token/domain"
)

type repository struct {
	db *sql.DB
}

func NewUserRepository(con *sql.DB) domain.UserRepository {
	return &repository{db: con}
}

// FindByID implements domain.UserRepository.
func (r repository) FindByID(ctx context.Context, id int64) (user domain.User, err error) {
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`

	err = r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.UserName,
		&user.Password,
	)

	if err == sql.ErrNoRows {
		return user, nil
	}
	return
}

// FindByUsername implements domain.UserRepository.
func (r repository) FindByUsername(ctx context.Context, username string) (user domain.User, err error) {
	query := `SELECT * FROM users WHERE username = ? LIMIT 1`

	err = r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.FullName,
		&user.Phone,
		&user.UserName,
		&user.Password,
	)

	if err == sql.ErrNoRows {
		return user, nil
	}
	return
}
