package main

import (
	"github.com/itsLeonB/drex/internal/config"
	"github.com/itsLeonB/drex/internal/delivery/grpc"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	configs := config.Load()
	s := grpc.Setup(configs)
	s.Run()
}
