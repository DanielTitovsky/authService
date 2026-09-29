package http_auth_routers

import (
	usecases "authService/internal/useCases"
	transport_http_utils "authService/internal/utils/transport/http"
	"fmt"
	"net/http"
	"time"
)

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responceHandler := transport_http_utils.NewHTTPResponceHandler(w)

	refreshToken, err := r.Cookie("refresh_token")

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Falide to parse coockie", "Wrong or undefined coockie")
		return
	}

	newAccessToken, newRefreshToken, err := h.service.RefreshToken(ctx, refreshToken.Value)

	if err != nil {
		fmt.Print(err)
		responceHandler.Error(http.StatusBadRequest, "Falide to refresh token", "Falide to refresh token")
		return
	}

	coockie := []*http.Cookie{
		{
			Name:     usecases.AccessTokenKey,
			Path:     "/",
			Value:    newAccessToken,
			Expires:  time.Now().Add(15 * time.Minute),
			HttpOnly: true,
		},
		{
			Name:     usecases.RefreshTokenKey,
			Path:     "/",
			Value:    newRefreshToken,
			Expires:  time.Now().Add(720 * time.Hour),
			HttpOnly: true,
		},
	}

	responceHandler.SetCookie(coockie)
	responceHandler.Success(http.StatusOK, true)
}
