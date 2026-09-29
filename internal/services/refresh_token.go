package auth_services

import (
	jwt_utils "authService/internal/utils/jwt"
	"context"
	"fmt"
)

func (s *AuthService) RefreshToken(ctx context.Context, tokenString string) (string, string, error) {

	userToken, err := jwt_utils.ParseAndVerify(tokenString)
	tokenHash := jwt_utils.HashedTokenString(tokenString)

	if err != nil {
		return "", "", fmt.Errorf("Invalid token: %w", err)
	}

	_, err = s.repo.GetRefteshToken(ctx, tokenHash)

	if err != nil {
		return "", "", fmt.Errorf("Invalid token: %w", err)
	}

	err = s.repo.RemoveRefreshToken(ctx, tokenHash)

	if err != nil {
		return "", "", fmt.Errorf("Failed to remove old refresh token: %w", err)
	}

	accesstoken, refreshtoken, err := s.createTokenPair(ctx, userToken.UserId)

	if err != nil {
		return "", "", fmt.Errorf("Failed to create token pair: %w", err)
	}

	return accesstoken, refreshtoken, nil
}
