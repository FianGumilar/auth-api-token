package domain

import (
	"context"

	"github.com/FianGumilar/auth-api-token/dto"
)

type User struct {
	ID       int64  `db:"id"`
	FullName string `db:"full_name"`
	Phone    string `db:"phone"`
	UserName string `db:"user_name"`
	Password string `db:"password"`
}

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
}

type UserService interface {
	Authenticate(ctx context.Context, req dto.AutRequest) (dto.AuthResponse, error)
	ValidateToken(ctx context.Context, token string) (dto.UserData, error)
}
