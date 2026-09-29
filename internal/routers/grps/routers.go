package grps_auth_routers

import (
	"github.com/DanielTitovsky/authgrpc"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authgrpc.UnimplementedAuthGRPCServer
	validator *validator.Validate
}

func NewAuthHandler(validator *validator.Validate) AuthHandler {
	return AuthHandler{
		validator: validator,
	}
}
