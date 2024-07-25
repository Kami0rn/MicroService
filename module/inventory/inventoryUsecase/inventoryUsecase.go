package inventoryUsecase

import "github.com/Kami0rn/MicroService/module/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase (inventoryRepository inventoryRepository.InventoryRepositoryService) InventoryUsecaseService{
	return &inventoryUsecase{
		inventoryRepository : inventoryRepository ,
	}
}