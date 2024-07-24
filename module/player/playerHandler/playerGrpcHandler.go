package playerHandler

import "github.com/Kami0rn/MicroService/module/player/playerUsecase"

type (
	playerGrpcHandlerService struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler (playerUsecase playerUsecase.PlayerUsecaseService) playerGrpcHandlerService{
	return playerGrpcHandlerService{playerUsecase}
}