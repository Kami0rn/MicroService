package main

import (
	"context"
	"log"
	"os"

	"github.com/Kami0rn/MicroService/config"
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

	log.Println(cfg)
}