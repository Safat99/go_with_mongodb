package deletion

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (
	result *mongo.DeleteResult, err error) {

	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteOne(ctx, query)
	return
}

func DeleteMany(client *mongo.Client, ctx context.Context, dataBase, col string, query interface{}) (
	result *mongo.DeleteResult, err error) {

	collection := client.Database(dataBase).Collection(col)
	result, err = collection.DeleteMany(ctx, query)
	return
}
