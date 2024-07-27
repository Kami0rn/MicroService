package server

import (
	"github.com/Kami0rn/MicroService/module/inventory/inventoryHandler"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryRepository"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	inventoryRepository := inventoryRepository.NewInventoryRepository(s.db)
	inventoryUsecase := inventoryUsecase.NewInventoryUsecase(inventoryRepository)
	inventoryHttpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, inventoryUsecase)
	inventoryGrpHandler := inventoryHandler.NewInventoryGrpcHandler(inventoryUsecase)
	inventoryQueue := inventoryHandler.NewInventoryQueueHandler(s.cfg, inventoryUsecase)

	_ = inventoryHttpHandler
	_ = inventoryGrpHandler
	_ = inventoryQueue

	inventory := s.app.Group("/inventory_v1")

	//Health Check
	inventory.GET("", s.healthCheckService)

}
