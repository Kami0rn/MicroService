package inventoryHandler

import (
	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler (cfg *config.Config , inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler {
		cfg : cfg,
		inventoryUsecase : inventoryUsecase ,
	}
}