package mongodb

import (
	"fmt"
	"os"

	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection

// var ConnectionURI = "mongodb://192.168.5.72:2701/"

func MongoDB() *mongo.Database {
	// fmt.Println(connectionURI)
	connectionURI := os.Getenv("MONGODB")
	clientOptions := options.Client().ApplyURI(connectionURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	conn := client.Database("accessLog")
	fmt.Println("MongoDB Connected")
	return conn
}
