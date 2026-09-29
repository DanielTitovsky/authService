package pgx_auth_repository

import "github.com/jackc/pgx/v5/pgxpool"

type AuthRepository struct {
	*pgxpool.Pool
}

func NewAuthRepository(pgxPool *pgxpool.Pool) AuthRepository {
	return AuthRepository{
		pgxPool,
	}
}
