package server

import (
	"log"
	playerPb "github.com/Kami0rn/MicroService/module/player/playerPb"
	"github.com/Kami0rn/MicroService/module/player/playerHandler"
	"github.com/Kami0rn/MicroService/module/player/playerRepository"
	"github.com/Kami0rn/MicroService/module/player/playerUsecase"
	"github.com/Kami0rn/MicroService/pkg/grpccon"
)

func (s *server) playerService() {
	playerRepository := playerRepository.NewPlayerRepository(s.db)
	playerUsecase := playerUsecase.NewPlayerUsecase(playerRepository)
	playerHttpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, playerUsecase)
	playerGrpcHandler := playerHandler.NewPlayerGrpcHandler(playerUsecase)
	playerQueue := playerHandler.NewPlayerQueueHandler(s.cfg , playerUsecase)

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, playerGrpcHandler)

		log.Printf("Player gRPC server listenign on %s", s.cfg.Grpc.PlayerUrl)

		grpcServer.Serve(lis)
	}()

	_ = playerHttpHandler
	_ = playerGrpcHandler
	_ = playerQueue

	player := s.app.Group("/player_v1")

	//Health Check
	player.GET("" , s.healthCheckService)

}