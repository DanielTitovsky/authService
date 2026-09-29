package auth_services

import (
	jwt_utils "authService/internal/utils/jwt"
	"context"
	"fmt"
)

func (s *AuthService) Logout(ctx context.Context, tokenString string) error {
	tokenHash := jwt_utils.HashedTokenString(tokenString)

	err := s.repo.RemoveRefreshToken(ctx, tokenHash)

	if err != nil {
		return fmt.Errorf("Failed to logout: %w", err)
	}

	return nil
}
