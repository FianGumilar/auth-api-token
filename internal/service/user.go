package service

import (
	"context"
	"encoding/json"

	"github.com/FianGumilar/auth-api-token/domain"
	"github.com/FianGumilar/auth-api-token/dto"
	"github.com/FianGumilar/auth-api-token/utils"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository  domain.UserRepository
	cacheRepository domain.CacheRepository
}

func NewUserService(userRepository domain.UserRepository, cacheRepository domain.CacheRepository) domain.UserService {
	return &service{
		userRepository:  userRepository,
		cacheRepository: cacheRepository,
	}
}

// Authenticate implements domain.UserService.
func (s service) Authenticate(ctx context.Context, req dto.AutRequest) (dto.AuthResponse, error) {
	// Check user exist
	user, err := s.userRepository.FindByUsername(ctx, req.Username)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	if user == (domain.User{}) {
		return dto.AuthResponse{}, domain.ErrAuthFailed
	}

	// Compare Hash and Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return dto.AuthResponse{}, nil
	}

	// Generate random string
	token := utils.GenerateRandmoString(16)

	userJson, _ := json.Marshal(user)
	_ = s.cacheRepository.Set("token : "+token, userJson)

	return dto.AuthResponse{
		Token: token,
	}, nil
}

// ValidateToken implements domain.UserService.
func (s service) ValidateToken(ctx context.Context, token string) (dto.UserData, error) {
	data, err := s.cacheRepository.Get("token : " + token)
	if err != nil {
		return dto.UserData{}, domain.ErrAuthFailed
	}

	var user domain.User
	_ = json.Unmarshal(data, &user)

	return dto.UserData{
		ID:       user.ID,
		FullName: user.FullName,
		Phone:    user.Phone,
		UserName: user.UserName,
	}, nil
}
