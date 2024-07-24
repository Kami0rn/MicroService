package playerUsecase

import (
	"github.com/Kami0rn/MicroService/module/player/playerRepository"
)

type (
	PlayerUsecaseService interface{}

	playerUsecase struct{
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase (playerRepository playerRepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase {playerRepository : playerRepository}
}