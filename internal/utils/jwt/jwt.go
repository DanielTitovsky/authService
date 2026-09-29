package jwt_utils

import (
	usecases "authService/internal/useCases"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Можно добавить дженерик который и принимать его в функции чтобы самому указывать какой дожен быть Claim так будет универсальнее
func GenerateToken(userId uuid.UUID, lifeTime time.Duration) (string, error) {
	claims := usecases.DefaultJwtClaim{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(lifeTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "DanDan",
			Subject:   "you",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(usecases.SecretKey))

	if err != nil {
		return "", fmt.Errorf("failed to signing token: %w", err)
	}

	return ss, nil
}

// тоже самое как будто можно переписать на дженерики и возвращать тогда можно будет нормальный обьект а не срань какую то
func ParseAndVerify(tokenString string) (usecases.DefaultJwtClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&usecases.DefaultJwtClaim{},
		func(token *jwt.Token) (any, error) {
			return []byte(usecases.SecretKey), nil
		},
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
		jwt.WithExpirationRequired(),
	)

	if err != nil {
		return usecases.DefaultJwtClaim{}, fmt.Errorf("Failed to parse token: %w", err)
	} else if claims, ok := token.Claims.(*usecases.DefaultJwtClaim); ok {
		return *claims, nil
	} else {
		return usecases.DefaultJwtClaim{}, fmt.Errorf("Unknow claim type, cannot proceed")
	}
}

func HashedTokenString(tokenString string) string {
	h := sha256.Sum256([]byte(tokenString))
	return hex.EncodeToString(h[:])
}
