package playerHandler

import (
	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
