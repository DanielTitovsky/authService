package pgx_auth_repository

import (
	"authService/internal/domain"
	"context"
	"fmt"
)

func (r *AuthRepository) RegisterUser(ctx context.Context, user domain.User) (domain.User, error) {
	var newUser domain.User

	sql := `
		INSERT INTO users (email, password_hash, username, display_name)
		VALUES ($1, $2, $3, $4)
		RETURNING id, password_hash, email, username, display_name;
	`

	err := r.Pool.QueryRow(ctx, sql,
		user.Email,
		user.PasswordHash,
		user.UserName,
		user.DisplayName,
	).Scan(
		&newUser.Id,
		&newUser.PasswordHash,
		&newUser.Email,
		&newUser.UserName,
		&newUser.DisplayName,
	)

	if err != nil {
		return domain.User{}, fmt.Errorf("Failed to save user: %w", err)
	}

	return newUser, nil
}
