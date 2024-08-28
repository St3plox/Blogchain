// file: testutil/mongo_container.go
package testutil

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestEnv struct {
	mongoContainer testcontainers.Container
	MongoClient    *mongo.Client
}

// SetupMongoDBContainer sets up the MongoDB container and returns a TestEnv.
func SetupMongoDBContainer(ctx context.Context) (*TestEnv, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:5.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}
	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start MongoDB container: %w", err)
	}

	host, err := mongoContainer.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get MongoDB container host: %w", err)
	}

	port, err := mongoContainer.MappedPort(ctx, "27017")
	if err != nil {
		return nil, fmt.Errorf("failed to get MongoDB container port: %w", err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	return &TestEnv{
		mongoContainer: mongoContainer,
		MongoClient:    client,
	}, nil
}

// Teardown stops the MongoDB container and cleans up resources.
func (env *TestEnv) Teardown(ctx context.Context) error {
	if err := env.MongoClient.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting MongoDB client: %v", err)
		return err
	}

	if err := env.mongoContainer.Terminate(ctx); err != nil {
		log.Printf("Error stopping MongoDB container: %v", err)
		return err
	}

	return nil
}
