package usecases

import "github.com/google/uuid"

type UserResponceDto struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	UserName string    `json:"username"`
}
