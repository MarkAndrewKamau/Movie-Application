package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // <-- this was missing
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// DBInstance creates and returns a connection (client) to MongoDB
func DBInstance() *mongo.Client {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to find .env file")
	}

	// Read the MongoDB URI from the environment
	mongoURI := os.Getenv("MONGODB_URI")

	// If the variable is empty, stop the program
	if mongoURI == "" {
		log.Fatal("MONGODB_URI not set!")
	}

	fmt.Println("MongoDB URI:", mongoURI)

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	return client
}

// Global variable for MongoDB client
var Client *mongo.Client = DBInstance()

// openCollection returns a collection object from the connected database
func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to find .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	fmt.Println("DATABASE_NAME:", databaseName)

	collection := Client.Database(databaseName).Collection(collectionName)

	return collection
}
