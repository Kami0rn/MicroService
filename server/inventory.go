package server

import (
	"log"

	"github.com/Kami0rn/MicroService/module/inventory/inventoryHandler"
	inventoryPb "github.com/Kami0rn/MicroService/module/inventory/inventoryPb"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryRepository"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"
	"github.com/Kami0rn/MicroService/pkg/grpccon"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryUsecase)
	inventoryQueue := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, inventoryGrpHandler)

		log.Printf("Inventory gRPC server listenign on %s", s.cfg.Grpc.InventoryUrl)

		grpcServer.Serve(lis)
	}()

	_ = inventoryHttpHandler
	_ = inventoryGrpHandler
	_ = inventoryQueue

	inventory := s.app.Group("/inventory_v1")

	//Health Check
	inventory.GET("", s.healthCheckService)

}
