package Models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var DB *mongo.Client //DB Client
func InitDB(mongoURI string) { //getting db object
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(mongoURI)
	fmt.Println("Init DB")
	db, err := mongo.Connect(ctx, clientOptions)
	DB = db
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = DB.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Connected to MongoDB!")

}