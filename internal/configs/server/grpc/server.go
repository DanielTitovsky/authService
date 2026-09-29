package app_grpc_server

import (
	"context"
	"net"
	"time"

	"github.com/DanielTitovsky/authgrpc"
	"google.golang.org/grpc"
)

type AuthGRPCServer struct {
	config     Config
	grpcServer *grpc.Server
}

func NewAuthGRPCServer(config Config, authHandler authgrpc.AuthGRPCServer) *AuthGRPCServer {

	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(config.connectionTimeout),
	)

	authgrpc.RegisterAuthGRPCServer(
		grpcServer,
		authHandler,
	)

	return &AuthGRPCServer{
		config:     config,
		grpcServer: grpcServer,
	}
}

func (a *AuthGRPCServer) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", ":"+a.config.port)

	if err != nil {
		return err
	}

	ch := make(chan error, 1)

	go func() {
		ch <- a.grpcServer.Serve(lis)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		stopped := make(chan struct{})

		go func() {
			a.grpcServer.GracefulStop()
			close(stopped)
		}()

		timer := time.NewTimer(a.config.shutDownTimeout)
		defer timer.Stop()

		select {
		case <-stopped:
			return nil

		case <-timer.C:
			a.grpcServer.Stop()
			<-stopped
			return nil
		}
	}
}
