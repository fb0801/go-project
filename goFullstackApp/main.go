package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID int `json:"_id" bson:"_id"`
	Completed bool `json:"completed"`
	Body string `json:"body"`
}

var collection *mongo.Collection

func main(){
	fmt.Println("hello world")

	err:= godotenv.Load(".env")
	if err != nil{
		log.Fatal("Error loading .env file", err)
	}

	MONGODB_URI:= os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client,err := mongo.Connect(context.Background(), clientOptions)

	if err != nil{
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	collection = client.Database("goland_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Get("/api/todos", createTodos)
	app.Get("/api/todos/:id", updateTodos)
	app.Get("/api/todos/:id", deleteTodos)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Fatal(app.Listen("0.0.0.0" + port))
	
}