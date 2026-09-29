package http_auth_routers

import (
	usecases "authService/internal/useCases"
	transport_http_utils "authService/internal/utils/transport/http"
	"net/http"
	"time"
)

func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responceHandler := transport_http_utils.NewHTTPResponceHandler(w)

	userData, err := transport_http_utils.GetAndFilterJson[usecases.RegisterTransportDto](h.validator, r)

	if err != nil {
		responceHandler.Error(400, err.Error(), "Invalid data format")
		return
	}

	registerResult, err := h.service.RegisterUser(ctx, usecases.RegisterServiceDto{
		Email:       userData.Email,
		Password:    userData.Password,
		UserName:    userData.UserName,
		DisplayName: userData.DisplayName,
	})

	if err != nil {
		responceHandler.Error(400, err.Error(), "Failed to register. \n Try again later")
		return
	}

	userResponceData := usecases.UserResponceDto{
		Id:       registerResult.User.Id,
		Email:    registerResult.User.Email,
		UserName: registerResult.User.UserName,
	}

	coockie := []*http.Cookie{
		{
			Name:     usecases.AccessTokenKey,
			Path:     "/",
			Value:    registerResult.AccessToken,
			Expires:  time.Now().Add(15 * time.Minute),
			HttpOnly: true,
		},
		{
			Name:     usecases.RefreshTokenKey,
			Path:     "/",
			Value:    registerResult.RefreshToken,
			Expires:  time.Now().Add(720 * time.Hour),
			HttpOnly: true,
		},
	}

	responceHandler.SetCookie(coockie)
	responceHandler.Success(200, userResponceData)
}
