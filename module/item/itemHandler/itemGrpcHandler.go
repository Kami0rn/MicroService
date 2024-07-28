package itemHandler

import (
	"context"
	itemPb "github.com/Kami0rn/MicroService/module/item/itemPb"
	"github.com/Kami0rn/MicroService/module/item/itemUsecase"
)

type (
	itemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{itemUsecase: itemUsecase}
}

func (g *itemGrpcHandler) FindItemsInIds(ctx context.Context,req *itemPb.FindItemInIdsReq) (*itemPb.FindItemInIdsRes , error){
	return nil,nil
}