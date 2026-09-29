package http_auth_routers

import (
	usecases "authService/internal/useCases"
	transport_http_utils "authService/internal/utils/transport/http"
	"net/http"
	"time"
)

func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	responceHandler := transport_http_utils.NewHTTPResponceHandler(w)

	userRequest, err := transport_http_utils.GetAndFilterJson[usecases.LoginTransportDto](h.validator, r)

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Parse data failde", "Invalid user Data")
		return
	}

	loginResult, err := h.service.LoginUser(ctx, usecases.LoginServicetDto{
		Password: userRequest.Password,
		Email:    userRequest.Email,
	})

	if err != nil {
		responceHandler.Error(http.StatusBadRequest, "Wrong password or email", "Wrong password or email \n please try again")
		return
	}

	userResponceData := usecases.UserResponceDto{
		Id:       loginResult.User.Id,
		Email:    loginResult.User.Email,
		UserName: loginResult.User.UserName,
	}

	coockie := []*http.Cookie{
		{
			Name:     usecases.AccessTokenKey,
			Path:     "/",
			Value:    loginResult.AccessToken,
			Expires:  time.Now().Add(15 * time.Minute),
			HttpOnly: true,
		},
		{
			Name:     usecases.RefreshTokenKey,
			Path:     "/",
			Value:    loginResult.RefreshToken,
			Expires:  time.Now().Add(720 * time.Hour),
			HttpOnly: true,
		},
	}

	responceHandler.SetCookie(coockie)
	responceHandler.Success(http.StatusAccepted, userResponceData)
}
