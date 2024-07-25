package main

import (
	"context"
	"log"
	"os"

	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/pkg/database"
	"github.com/Kami0rn/MicroService/server"
)

func main() {
	ctx := context.Background()

	// Initialize Config 
	cfg := config.LoadConfig(func () string{
		if len(os.Args) < 2 {
			log.Fatal("Please provide a .env path config file")
		}
		return os.Args[1]
	}())

	//Database connection
	db :=  database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	//Start server
	server.Start(ctx , &cfg , db)

}