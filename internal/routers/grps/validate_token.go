package grps_auth_routers

import (
	jwt_utils "authService/internal/utils/jwt"
	"context"
	"fmt"

	"github.com/DanielTitovsky/authgrpc"
)

func (a *AuthHandler) ValidateToken(ctx context.Context, req *authgrpc.TokenString) (*authgrpc.ValidateTokenResponse, error) {
	tokenData, err := jwt_utils.ParseAndVerify(req.GetToken())

	if err != nil {
		return nil, fmt.Errorf("Invalid token: %w", err)
	}

	return &authgrpc.ValidateTokenResponse{
		UserId: tokenData.UserId.String(),
	}, nil
}
