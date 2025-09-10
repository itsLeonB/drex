package delivery_test

import (
	"testing"

	"github.com/itsLeonB/drex/internal/config"
	grpcDelivery "github.com/itsLeonB/drex/internal/delivery/grpc"
	"github.com/stretchr/testify/assert"
)

func TestGrpcSetup(t *testing.T) {
	// Skip this test as it requires actual database connection
	t.Skip("Skipping gRPC setup test as it requires database connection")

	configs := config.Config{
		App: config.App{
			Name: "TestApp",
			Env:  "debug",
			Port: "50051",
		},
		DB: config.DB{
			Driver:   "postgres",
			Host:     "localhost",
			Port:     "5432",
			User:     "test",
			Password: "test",
			Name:     "test",
		},
	}

	server := grpcDelivery.Setup(configs)
	assert.NotNil(t, server)
}
