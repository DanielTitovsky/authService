package pgx_auth_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *AuthRepository) GetRefteshToken(ctx context.Context, tokenHash string) (string, error) {
	var newTokenHash string

	sql := `
	SELECT
    	token_hash
	FROM public.refresh_tokens
	WHERE token_hash = $1
  		AND revoked_at IS NULL
  		AND expires_at > now();
	`

	err := r.Pool.QueryRow(ctx, sql,
		tokenHash,
	).Scan(&newTokenHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("Token has expired or invalid: %w", err)
		}
		return "", fmt.Errorf("Failed to getting token: %w", err)
	}

	return newTokenHash, nil
}
