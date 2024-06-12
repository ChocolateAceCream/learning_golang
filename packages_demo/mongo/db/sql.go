package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mg.Client
}

func (m *MongoDB) Init() error {
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	maxPoolSize := 20

	uri := fmt.Sprintf("mongodb://%s:%s@localhost:27017/?maxPoolSize=%v&w=majority", username, password, maxPoolSize)

	// use the stable api v1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mg.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	m.Client = client
	return nil
}

func (m *MongoDB) Disconnect() error {
	return m.Client.Disconnect(context.TODO())
}

func (m *MongoDB) Ping() (result bson.M, err error) {
	err = m.Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	return result, err
}
