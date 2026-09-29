package usecases

import "authService/internal/domain"

type LoginTransportDto struct {
	Email    string `json:"email" validate:"required,email" db:"email"`
	Password string `json:"password" validate:"required,min=8,max=20" db:"password_hash"`
}

type LoginServicetDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	User         domain.User
	AccessToken  string
	RefreshToken string
}
