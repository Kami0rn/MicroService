package server

import (
	"github.com/Kami0rn/MicroService/module/item/itemRepository"
	"github.com/Kami0rn/MicroService/module/item/itemUsecase"
	"github.com/Kami0rn/MicroService/module/item/itemHandler"
)


func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpchandler := itemHandler.NewItemGrpcHandler(itemUsecase)

	_ = itemHttpHandler
	_ = itemGrpchandler


	item := s.app.Group("/item_v1")

	//Health Check
	item.GET("" , s.healthCheckService)

}