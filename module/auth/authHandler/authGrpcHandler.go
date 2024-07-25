package authHandler

import "github.com/Kami0rn/MicroService/module/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthgrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{authUsecase}
}