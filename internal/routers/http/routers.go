package http_auth_routers

import (
	app_http_server "authService/internal/configs/server/http"
	"authService/internal/midelware"
	usecases "authService/internal/useCases"
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service   AuthService
	validator *validator.Validate
}

type AuthService interface {
	LoginUser(ctx context.Context, user usecases.LoginServicetDto) (usecases.LoginResult, error)
	RefreshToken(ctx context.Context, tokenString string) (string, string, error)
	Logout(ctx context.Context, tokenString string) error
	RegisterUser(ctx context.Context, user usecases.RegisterServiceDto) (usecases.RegisterResult, error)
}

func NewAuthHandler(service AuthService, validator *validator.Validate) AuthHandler {
	return AuthHandler{
		service:   service,
		validator: validator,
	}
}

func (h *AuthHandler) Routers() []app_http_server.Route {
	return []app_http_server.Route{
		{
			Path:    "/register",
			Method:  http.MethodPost,
			Handler: h.RegisterUser,
		},
		{
			Path:    "/login",
			Method:  http.MethodPost,
			Handler: h.LoginUser,
		},
		{
			Path:    "/logout",
			Method:  http.MethodPost,
			Handler: h.Logout,
			Middlewares: []midelware.Middleware{
				midelware.AuthMiddleware(),
			},
		},
		{
			Path:    "/refresh",
			Method:  http.MethodPost,
			Handler: h.Refresh,
			Middlewares: []midelware.Middleware{
				midelware.AuthMiddleware(),
			},
		},
	}
}

// {
//   "email": "Test2@mail.ru",
//   "password": "TestTest2!",
//   "userName": "test22222",
//   "displayName": "test2222"
// }
