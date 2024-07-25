package server

import (
	"github.com/Kami0rn/MicroService/module/auth/authHandler"
	"github.com/Kami0rn/MicroService/module/auth/authRepository"
	"github.com/Kami0rn/MicroService/module/auth/authUsecase"
)

func (s *server) authService() {
	authRepository := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepository)
	authHttpHandler := authHandler.NewAuthHttphHandler(s.cfg ,authUsecase)
	authGrpHandler := authHandler.NewAuthGrpcHandler(authUsecase)

	_ = authHttpHandler
	_ = authGrpHandler

	auth := s.app.Group("/auth_v1")

	//Health Check
	auth.GET("" , s.healthCheckService)

}