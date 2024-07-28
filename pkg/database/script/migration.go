package main

import (
	"context"
	"log"
	"os"

	"github.com/Kami0rn/MicroService/config"
	"github.com/Kami0rn/MicroService/pkg/database/migration"
)

func main() {
	ctx := context.Background()
	_ = ctx
	
	// Initialize Config 
	cfg := config.LoadConfig(func () string{
		if len(os.Args) < 2 {
			log.Fatal("Please provide a .env path config file")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "player" :
	case "auth" : 
		migration.AuthMigrate(ctx , &cfg)
	case "item" :
	case "inventory" :
	case "payment" :
	}


}