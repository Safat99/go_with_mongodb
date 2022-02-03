package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	// ctx will be used to set deadline for process, here deadline will of 30 secs
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err

}

// select database and collection with Client.Database method
// and Database.Collection method
func insertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func insertMany(client *mongo.Client, ctx context.Context, dataBase, col string, docs []interface{}) (*mongo.InsertManyResult, error) {
	collection := client.Database(dataBase).Collection(col)

	result, err := collection.InsertMany(ctx, docs)
	return result, err
}

func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected succesfully")
	return nil
}

func main() {
	client, ctx, cancel, err := connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)
	ping(client, ctx)

	// creating a object of type interface to store
	// the bson values, that we are inserting into database
	var document interface{}

	document = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	insertOneResult, err := insertOne(client, ctx, "gfg", "marks", document)

	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	var documents []interface{}

	documents = []interface{}{
		bson.D{
			{"rollNo", 153},
			{"maths", 65},
			{"science", 59},
			{"computer", 55},
		},
		bson.D{
			{"rollNo", 162},
			{"maths", 86},
			{"science", 80},
			{"computer", 69},
		},
	}

	insertManyResult, err := insertMany(client, ctx, "gfg", "marks", documents)

	if err != nil {
		panic(err)
	}

	fmt.Println("Result of InsertMany")

	for id := range insertManyResult.InsertedIDs {
		fmt.Println(id)
	}
}
