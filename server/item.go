package server

import (
	"log"
	itemPb "github.com/Kami0rn/MicroService/module/item/itemPb"
	"github.com/Kami0rn/MicroService/module/item/itemHandler"
	"github.com/Kami0rn/MicroService/module/item/itemRepository"
	"github.com/Kami0rn/MicroService/module/item/itemUsecase"
	"github.com/Kami0rn/MicroService/pkg/grpccon"
)


func (s *server) itemService() {
	itemRepository := itemRepository.NewItemRepository(s.db)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepository)
	itemHttpHandler := itemHandler.NewItemHttpHandler(s.cfg, itemUsecase)
	itemGrpcHandler := itemHandler.NewItemGrpcHandler(itemUsecase)

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, itemGrpcHandler)

		log.Printf("Item gRPC server listenign on %s", s.cfg.Grpc.ItemUrl)

		grpcServer.Serve(lis)
	}()

	_ = itemHttpHandler
	_ = itemGrpcHandler


	item := s.app.Group("/item_v1")

	//Health Check
	item.GET("" , s.healthCheckService)

}