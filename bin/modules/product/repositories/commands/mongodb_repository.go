package commands

import (
	"context"

	"codebase-go/bin/modules/product"
	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type commandMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewCommandMongodbRepository(mongodb mongodb.MongoDBLogger) product.MongodbRepositoryCommand {
	return &commandMongodbRepository{
		mongoDb: mongodb,
	}
}

func (c commandMongodbRepository) NewObjectID(ctx context.Context) string {
	return primitive.NewObjectID().Hex()
}

func (c commandMongodbRepository) InsertOneProduct(ctx context.Context, data models.Product) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)
		err := c.mongoDb.InsertOne(mongodb.InsertOne{
			CollectionName: "products",
			Document:       data,
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}
	}()

	return output
}

func (c commandMongodbRepository) UpdateOneProduct(ctx context.Context, data models.Product) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, err := primitive.ObjectIDFromHex(data.Id)
		if err != nil {
			output <- utils.Result{Error: err}
		}

		err = c.mongoDb.UpdateOne(mongodb.UpdateOne{
			CollectionName: "products",
			Document:       data.UpsertProduct(),
			Filter: bson.M{
				"id": objId,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}

		output <- utils.Result{Data: data}
	}()

	return output
}

func (c commandMongodbRepository) DeleteOneProduct(ctx context.Context, productId string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, err := primitive.ObjectIDFromHex(productId)
		if err != nil {
			output <- utils.Result{Error: err}
		}

		err = c.mongoDb.DeleteOne(mongodb.DeleteOne{
			CollectionName: "product",
			Filter: bson.M{
				"id": objId,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{Error: err}
		}
	}()

	return output
}
