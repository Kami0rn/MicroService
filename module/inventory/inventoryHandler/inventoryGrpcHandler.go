package inventoryHandler

import (
	"context"
	inventoryPb "github.com/Kami0rn/MicroService/module/inventory/inventoryPb"
	"github.com/Kami0rn/MicroService/module/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
		inventoryPb.UnimplementedInventoryGrpcServiceServer
	}
)

func NewInventoryGrpcHandler (inventoryUsecase inventoryUsecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase : inventoryUsecase,
	}
}

func (g *inventoryGrpcHandler) IsAvailableToSell (ctx context.Context, req *inventoryPb.IsAvailableToSellReq)(*inventoryPb.IsAvailableToSellRes , error){
	return nil,nil
}