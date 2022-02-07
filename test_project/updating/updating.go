package updating

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (
	result *mongo.UpdateResult, err error) {

	// select the database and the collection/table
	collection := client.Database(dataBase).Collection(col)

	// if the document and the filter matches, update the documents
	// update contains the field(column) which is going to be updated
	result, err = collection.UpdateOne(ctx, filter, update)
	return
}

func UpdateMany(client *mongo.Client, ctx context.Context, dataBase, col string, filter, update interface{}) (
	result *mongo.UpdateResult, err error) {

	collection := client.Database(dataBase).Collection(col)
	result, err = collection.UpdateMany(ctx, filter, update)
	return
}
