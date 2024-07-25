package inventoryHandler

import "github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler (inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase : inventoryUsecase,
	}
}