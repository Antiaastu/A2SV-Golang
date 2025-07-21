package main

import (
	"Task6/data"
	"Task6/router"
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("❌ MONGODB_URI environment variable not set")
	}
	
	err := data.InitMongoDB(ctx, mongoURI)
	if err != nil{
		log.Fatal("Failed to connect to MongoDB", err)
	}

	defer data.CloseMongoDB()

	data.InitUserCollection()
	r := router.SetUpRouter()
	r.Run(":8080")
}