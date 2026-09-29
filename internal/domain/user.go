package domain

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uuid.UUID `json:"id" db:"id"`
	Email        string    `json:"email" validate:"required,email" db:"email"`
	PasswordHash string    `json:"passwordHash" validate:"required,min=8,max=20" db:"password_hash"`
	UserName     string    `json:"userName" validate:"required,min=8,max=64" db:"user_name"`
	DisplayName  string    `json:"displayName" validate:"required,min=8,max=64" db:"display_name"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) HashingPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytePassword), err
}
