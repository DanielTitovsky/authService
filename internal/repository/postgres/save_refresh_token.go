package pgx_auth_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (r *AuthRepository) SaveRefreshToken(ctx context.Context, UserId uuid.UUID, ExpairedAt time.Time, TokenHash string) error {
	sql := `
	INSERT INTO public.refresh_tokens (
    	user_id,
    	token_hash,
    	expires_at
	)
	VALUES ($1, $2, $3)
	`

	_, err := r.Pool.Exec(ctx, sql,
		UserId,
		TokenHash,
		ExpairedAt,
	)

	if err != nil {
		return fmt.Errorf("Failed to saving token: %w", err)
	}

	return nil
}
