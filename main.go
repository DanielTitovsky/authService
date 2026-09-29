package main

import (
	app_repository "authService/internal/configs/repository"
	app_grpc_server "authService/internal/configs/server/grpc"
	app_http_server "authService/internal/configs/server/http"
	"authService/internal/midelware"
	pgx_repository "authService/internal/repository/postgres"
	grps_auth_routers "authService/internal/routers/grps"
	http_auth_routers "authService/internal/routers/http"
	auth_services "authService/internal/services"
	"context"
	"fmt"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"

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

	//Создание сервера HTTP
	httpServerConfig := app_http_server.NewHttpServerConfig()
	httpServer := app_http_server.NewServer(httpServerConfig, midelware.RequestId())
	router := app_http_server.NewRouter("v01", httpServer.ServerMux)

	//Создание бд
	pgxConnectinPoolConfig := app_repository.NewConfig()
	connectinPool := app_repository.NewPgxConnectinPool(ctx, pgxConnectinPoolConfig)

	validate := validator.New(validator.WithRequiredStructEnabled())

	//Auth
	authRepository := pgx_repository.AuthRepository(connectinPool)
	authService := auth_services.NewAuthService(&authRepository)
	authHandlers := http_auth_routers.NewAuthHandler(&authService, validate)

	//Создание сервера GRPC
	grpcHandler := grps_auth_routers.NewAuthHandler(validate)
	grpcConfig := app_grpc_server.NewConfig("50051", 5*time.Second, 15*time.Second)
	grpcServer := app_grpc_server.NewAuthGRPCServer(grpcConfig, &grpcHandler)

	router.RegisterRouters(authHandlers.Routers()...)

	errCh := make(chan error, 2)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		fmt.Print("http server started")
		errCh <- httpServer.Run(ctx)
	}()

	go func() {
		defer wg.Done()
		fmt.Print("\n")
		fmt.Print("grpc server started")
		errCh <- grpcServer.Run(ctx)
	}()

	select {
	case err := <-errCh:
		log.Printf("Server error: %v", err)
		cancel()

	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	wg.Wait()
}
