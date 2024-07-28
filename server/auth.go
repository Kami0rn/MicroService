package server

import (
	"log"

	"github.com/Kami0rn/MicroService/module/auth/authHandler"
	authPb "github.com/Kami0rn/MicroService/module/auth/authPb"
	"github.com/Kami0rn/MicroService/module/auth/authRepository"
	"github.com/Kami0rn/MicroService/module/auth/authUsecase"
	"github.com/Kami0rn/MicroService/pkg/grpccon"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	authHttpHandler := authHandler.NewAuthHttphHandler(s.cfg ,authUsecase)
	authGrpHandler := authHandler.NewAuthGrpcHandler(authUsecase)

	//gRPC 
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, authGrpHandler)

		log.Printf("Auth gRPC server listenign on %s", s.cfg.Grpc.AuthUrl)

		grpcServer.Serve(lis)
	}()

	_ = authHttpHandler
	_ = authGrpHandler

	auth := s.app.Group("/auth_v1")

	//Health Check
	auth.GET("" , s.healthCheckService)

}