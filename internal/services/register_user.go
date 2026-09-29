package auth_services

import (
	"authService/internal/domain"
	usecases "authService/internal/useCases"
	"context"
	"fmt"
)

func (s *AuthService) RegisterUser(ctx context.Context, userData usecases.RegisterServiceDto) (usecases.RegisterResult, error) {
	var err error
	user := domain.User{
		Email:       userData.Email,
		UserName:    userData.UserName,
		DisplayName: userData.DisplayName,
	}

	user.PasswordHash, err = user.HashingPassword(userData.Password)

	if err != nil {
		return usecases.RegisterResult{}, fmt.Errorf("Failed hashing password: %w", err)
	}

	savedUser, err := s.repo.RegisterUser(ctx, user)

	if err != nil {
		return usecases.RegisterResult{}, fmt.Errorf("Failed saving user: %w", err)
	}

	accessToken, refreshToken, err := s.createTokenPair(ctx, savedUser.Id)

	if err != nil {
		return usecases.RegisterResult{}, fmt.Errorf("Failed to create token pair: %w", err)
	}

	return usecases.RegisterResult{
		User:         savedUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
