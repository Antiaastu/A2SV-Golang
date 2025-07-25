package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("Error loading .env file: ", err)
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("DB connection error", err)
	}
	dbName := os.Getenv("DB_NAME")
	DB = client.Database(dbName)
	fmt.Println("Connected to the DB")
}
