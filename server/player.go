package server

import (
	"github.com/Kami0rn/MicroService/module/player/playerRepository"
	"github.com/Kami0rn/MicroService/module/player/playerUsecase"
	"github.com/Kami0rn/MicroService/module/player/playerHandler"
)

func (s *server) playerService() {
	playerRepository := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpHandler := playerHandler.NewPlayerGrpcHandler(playerUsecase)
	playerQueue := playerHandler.NewPlayerQueueHandler(s.cfg , playerUsecase)

	_ = playerHttpHandler
	_ = playerGrpHandler
	_ = playerQueue

	player := s.app.Group("/player_v1")

	//Health Check
	player.GET("" , s.healthCheckService)

}