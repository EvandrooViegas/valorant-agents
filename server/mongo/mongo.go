package mongo

import (
	"context"
	"fmt"
	"os"
	"valorant-agents/utils"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

func ConnectToMongoDB() *mongo.Client {
	// load .env file
	err := godotenv.Load("./env")
	utils.HandleErr(err)

	// get the MONGODB_URI from .env
	MONGODB_URI := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MONGODB_URI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	MongoDB, err = mongo.Connect(context.TODO(), opts)
	utils.HandleErr(err)

	// Send a ping to confirm a successful connection
	if err := MongoDB.Database("admin").RunCommand(context.TODO(), bson.D{{ "ping", 1 }}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("You successfully connected to MongoDB 📦")
	return MongoDB
}

func DisconnectFromMongoDB() {
	if err := MongoDB.Disconnect(context.TODO()); err != nil {
		panic(err)
	} else {
		fmt.Println("You successfully disconnected from MongoD 📦")
	}
}