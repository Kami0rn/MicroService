package authHandler

import (
	"context"

	authPb "github.com/Kami0rn/MicroService/module/auth/authPb"
	"github.com/Kami0rn/MicroService/module/auth/authUsecase"
)


type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{
		authUsecase : authUsecase,
	}
}

func (g *authGrpcHandler) CredentialSearch(ctx context.Context ,req *authPb.AccessTokenSearchReq)(*authPb.AccessTokenSearchRes, error) {
	return nil,nil
}

func (g *authGrpcHandler) RoleCount(ctx context.Context ,req *authPb.RoleCountReq)(*authPb.RoleCountRes, error) {
	return nil,nil
}