package authHandler

import "github.com/Kami0rn/MicroService/module/auth/authUsecase"

type (
	authgrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthgrpcHandler(authUsecase authUsecase.AuthUsecaseService) authUsecase.AuthUsecaseService{
	return &authgrpcHandler{authUsecase}
}