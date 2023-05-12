package queries

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"codebase-go/bin/modules/product"
	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/databases/mongodb"
	"codebase-go/bin/pkg/utils"
)

type queryMongodbRepository struct {
	mongoDb mongodb.MongoDBLogger
}

func NewQueryMongodbRepository(mongodb mongodb.MongoDBLogger) product.MongodbRepositoryQuery {
	return &queryMongodbRepository{
		mongoDb: mongodb,
	}
}

func (q queryMongodbRepository) FindOne(ctx context.Context, productId string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		objId, _ := primitive.ObjectIDFromHex(productId)
		var product models.Product
		err := q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &product,
			CollectionName: "product",
			Filter: bson.M{
				"id": objId,
			},
		}, ctx)
		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: product,
		}

	}()

	return output
}

func (q queryMongodbRepository) FindOneByUsername(ctx context.Context, Productname string) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var product models.Product
		err := q.mongoDb.FindOne(mongodb.FindOne{
			Result:         &product,
			CollectionName: "products",
			Filter: bson.M{
				"productname": Productname,
			},
		}, ctx)

		if err != nil {
			output <- utils.Result{
				Error: err,
			}
		}

		output <- utils.Result{
			Data: product,
		}

	}()

	return output
}
