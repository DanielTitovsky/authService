package auth_services

import (
	usecases "authService/internal/useCases"
	"context"
	"errors"
	"fmt"
)

func (s *AuthService) LoginUser(ctx context.Context, userData usecases.LoginServicetDto) (usecases.LoginResult, error) {
	user, err := s.repo.GetUserByEmail(ctx, userData.Email)

	if err != nil {
		return usecases.LoginResult{}, fmt.Errorf("Failded to getting user by email: %w", err)
	}

	if isValidPassword := user.ValidatePassword(userData.Password); !isValidPassword {
		return usecases.LoginResult{}, errors.New("wrong password")
	}

	accessToken, refreshToken, err := s.createTokenPair(ctx, user.Id)

	if err != nil {
		return usecases.LoginResult{}, fmt.Errorf("Failed to create token pair: %w", err)
	}

	return usecases.LoginResult{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
