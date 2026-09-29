package pgx_auth_repository

import (
	"authService/internal/domain"
	"context"
	"fmt"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	sql := `
	SELECT id, email, username, display_name, password_hash
	FROM users
	WHERE email = $1
	`

	err := r.Pool.QueryRow(ctx, sql,
		&email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.UserName,
		&user.DisplayName,
		&user.PasswordHash,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("Failed to getting user: %w", err)
	}

	return user, nil
}
