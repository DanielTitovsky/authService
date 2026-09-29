package auth_services

import (
	usecases "authService/internal/useCases"
	jwt_utils "authService/internal/utils/jwt"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *AuthService) createTokenPair(ctx context.Context, userId uuid.UUID) (string, string, error) {
	accessToken, err := jwt_utils.GenerateToken(userId, usecases.AccessTokenLifetime*time.Minute)

	if err != nil {
		return "", "", fmt.Errorf("Failed to generate access token: %w", err)
	}

	refreshToken, err := jwt_utils.GenerateToken(userId, usecases.RefreshTokenLifetime*time.Hour)

	if err != nil {
		return "", "", fmt.Errorf("Failed to generate refresh token: %w", err)
	}

	err = s.repo.SaveRefreshToken(
		ctx,
		userId,
		time.Now().Add(usecases.RefreshTokenLifetime*time.Hour),
		jwt_utils.HashedTokenString(refreshToken),
	)

	if err != nil {
		return "", "", fmt.Errorf("Failed to save token: %w", err)
	}

	return accessToken, refreshToken, nil
}
