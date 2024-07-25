package inventoryHandler

import (
	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler (cfg *config.Config , inventoryUsecase inventoryUsecase.InventoryUsecaseService ) InventoryQueueHandlerService {
	return &inventoryQueueHandler {
		cfg : cfg ,
		inventoryUsecase : inventoryUsecase ,
	}
}