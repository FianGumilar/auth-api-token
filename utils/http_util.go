package utils

import (
	"errors"

	"github.com/FianGumilar/auth-api-token/domain"
)

func GetHttpStatus(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500
	}
}
