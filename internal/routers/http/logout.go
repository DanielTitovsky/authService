package http_auth_routers

import (
	usecases "authService/internal/useCases"
	transport_http_utils "authService/internal/utils/transport/http"
	"fmt"
	"net/http"
	"time"
)

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responceHandler := transport_http_utils.NewHTTPResponceHandler(w)

	refreshToken, err := r.Cookie("refresh_token")

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Falide to logout", "Refresh token not found")
		return
	}

	err = h.service.Logout(ctx, refreshToken.Value)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Falide to logout", fmt.Sprintf("Faild to logout: %w", err))
		return
	}

	coockie := []*http.Cookie{
		{
			Name:     usecases.RefreshTokenKey,
			Path:     "/",
			Value:    "",
			MaxAge:   -1,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		},
		{
			Name:     usecases.AccessTokenKey,
			Path:     "/",
			Value:    "",
			MaxAge:   -1,
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		},
	}

	responceHandler.SetCookie(coockie)
	responceHandler.Success(http.StatusOK, true)
}
