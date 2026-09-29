package pgx_auth_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) RemoveRefreshToken(ctx context.Context, tokenHash string) error {
	sql := `
	DELETE FROM public.refresh_tokens
	WHERE token_hash = $1;
	`

	_, err := r.Pool.Exec(ctx, sql,
		tokenHash,
	)

	if err != nil {
		return fmt.Errorf("Failed to saving token: %w", err)
	}

	return nil
}
