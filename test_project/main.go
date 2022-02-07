package main

import (
	"context"
	"fmt"
	"time"

	"test_project/updating"

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
	// var document interface{}

	// document = bson.D{
	// 	{"rollNo", 175},
	// 	{"maths", 80},
	// 	{"science", 90},
	// 	{"computer", 95},
	// }

	// insertOneResult, err := insert.InsertOne(client, ctx, "gfg", "marks", document)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Result of InsertOne")
	// fmt.Println(insertOneResult.InsertedID)

	// var documents []interface{}

	// documents = []interface{}{
	// 	bson.D{
	// 		{"rollNo", 153},
	// 		{"maths", 65},
	// 		{"science", 59},
	// 		{"computer", 55},
	// 	},
	// 	bson.D{
	// 		{"rollNo", 162},
	// 		{"maths", 86},
	// 		{"science", 80},
	// 		{"computer", 69},
	// 	},
	// }

	// insertManyResult, err := insert.InsertMany(client, ctx, "gfg", "marks", documents)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Result of InsertMany")

	// for id := range insertManyResult.InsertedIDs {
	// 	fmt.Println(id)
	// }

	///////////////////////////////////////////////// finding section ///////////////////////////////////
	// var filter, option interface{}

	// filter = bson.D{
	// 	{"maths", bson.D{{"$gt", 70}}},
	// }

	// option = bson.D{{"_id", 0}}

	// cursor, err := find.Query(client, ctx, "gfg", "marks", filter, option)

	// if err != nil {
	// 	panic(err)
	// }

	// var resutls []bson.D

	// if err := cursor.All(ctx, &resutls); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Query Result")
	// for _, doc := range resutls {
	// 	fmt.Println(doc)
	// }

	/////////////////////////////////////////// update values section //////////////////////////////////////////

	// the field/columns of the documents that need to update

	filter := bson.D{
		{"maths", bson.D{{"$lt", 100}}},
	}

	update := bson.D{
		{"$set", bson.D{
			{"maths", 100},
		}},
	}

	result, err := updating.UpdateMany(client, ctx, "gfg", "marks", filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("update multiple documents")
	fmt.Println(result.ModifiedCount)
}
