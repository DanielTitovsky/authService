package usecases

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type DefaultJwtClaim struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

const AccessTokenKey = "access_token"
const RefreshTokenKey = "refresh_token"
const SecretKey = "Thedull melancholy, and no one to offer a hand In times of a destitute feeling"
const AccessTokenLifetime = 15
const RefreshTokenLifetime = 43200
