package auth_services

import (
	"authService/internal/domain"
	usecases "authService/internal/useCases"
	"context"
	"time"

	"github.com/google/uuid"
)

type AuthService struct {
	repo        AuthRepository
	kafkaClient KafkaClient
}

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	RemoveRefreshToken(ctx context.Context, tokenHash string) error
	SaveRefreshToken(ctx context.Context, UserId uuid.UUID, ExpairedAt time.Time, tokenHash string) error
	GetRefteshToken(ctx context.Context, tokenHash string) (string, error)
}

type KafkaClient interface {
	PublishUserRegistered(ctx context.Context, event usecases.RegisterUserEvent) error
}

func NewAuthService(repo AuthRepository, kafkaClient KafkaClient) AuthService {
	return AuthService{
		repo:        repo,
		kafkaClient: kafkaClient,
	}
}
