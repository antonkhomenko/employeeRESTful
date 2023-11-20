package main

import (
	"context"
	"database-lesson/storage"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const Port = ":8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env file")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("URI")))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	router := gin.Default()
	//mongoDB := storage.NewMongoDB(client)
	memoryStorage := storage.NewMemoryStorage()

	handler := NewHandler(memoryStorage)
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.ReadEmployee)
	router.GET("/employee", handler.ReadAllEmployees)
	router.PUT("/employee/:id", handler.UpdateEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	log.Fatal(router.Run(Port))
}
