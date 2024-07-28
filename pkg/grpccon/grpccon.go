package grpccon

import (
	"errors"
	// "fmt"
	"log"
	"net"
	playerPb "github.com/Kami0rn/MicroService/module/player/playerPb"
	itemPb "github.com/Kami0rn/MicroService/module/item/itemPb"
	inventoryPb "github.com/Kami0rn/MicroService/module/inventory/inventoryPb"
	authPb "github.com/Kami0rn/MicroService/module/auth/authPb"
	"github.com/Kami0rn/MicroService/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
	grpcAuth struct {
	}
)

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}

func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler,error) {
	opts := make([]grpc.DialOption,0)

	opts = append(opts , grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: GRPC client connection failed: %s",err.Error())
		return nil,errors.New("error: GRPC client connection failed")
	}
	return &grpcClientFactory {
		client: clientConn,
	},nil
}

func NewGrpcServer(config *config.Jwt, host string) (*grpc.Server , net.Listener) {
	opts := make([]grpc.ServerOption,0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal("Error: Failed to listen GRPC : ",err)
	}
	return grpcServer , lis
}