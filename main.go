package main

import (
	app_repository "authService/internal/configs/repository"
	app_server "authService/internal/configs/server"
	"authService/internal/midelware"
	pgx_repository "authService/internal/repository/postgres"
	http_auth_routers "authService/internal/routers/http"
	auth_services "authService/internal/services"
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/go-playground/validator/v10"
)

func main() {

	//Создание контекста
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	//Создание сервера
	httpServerConfig := app_server.NewHttpServerConfig()
	httpServer := app_server.NewServer(httpServerConfig, midelware.RequestId())
	router := app_server.NewRouter("v01", httpServer.ServerMux)

	//Создание бд
	pgxConnectinPoolConfig := app_repository.NewConfig()
	connectinPool := app_repository.NewPgxConnectinPool(ctx, pgxConnectinPoolConfig)

	validate := validator.New(validator.WithRequiredStructEnabled())

	//Auth
	authRepository := pgx_repository.AuthRepository(connectinPool)
	authService := auth_services.NewAuthService(&authRepository)
	authHandlers := http_auth_routers.NewAuthHandler(&authService, validate)

	router.RegisterRouters(authHandlers.Routers()...)

	fmt.Print("server started")

	httpServer.Run(ctx)
}
