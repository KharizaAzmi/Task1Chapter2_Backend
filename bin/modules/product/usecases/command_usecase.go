package usecases

import (
	"codebase-go/bin/modules/product"
	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/utils"
	"context"
)

type commandUsecase struct {
	productRepositoryQuery   product.MongodbRepositoryQuery
	productRepositoryCommand product.MongodbRepositoryCommand
	redis                 redis.Collections
}

func NewCommandUsecase(mq product.MongodbRepositoryQuery, mc product.MongodbRepositoryCommand, rc redis.Collections) product.UsecaseCommand {
	return &commandUsecase{
		productRepositoryQuery:   mq,
		productRepositoryCommand: mc,
		redis:                 rc,
	}
}

func (c commandUsecase) CreateProduct(ctx context.Context, payload models.Product) utils.Result {
	var result utils.Result

	// queryRes := <-c.productRepositoryQuery.FindOneByProductname(ctx, payload.ProductName)
	// if queryRes.Data != nil {
	// 	errObj := errors.Conflict("Product already exist")
	// 	result.Error = errObj
	// 	return result
	// }

	//payload.ProductName = utils.HashPassword(payload.ProductName)

	result = <-c.productRepositoryCommand.InsertOneProduct(ctx, payload)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed insert user")
		result.Error = errObj
		return result
	}

	return result
}

func (c commandUsecase) UpdateProduct(ctx context.Context, payload models.Product) utils.Result {
	var result utils.Result

	queryRes := <-c.productRepositoryQuery.FindOne(ctx, payload.Id)
	if queryRes.Data == nil {
		errObj := errors.NotFound("Product not found")
		result.Error = errObj
		return result
	}

	//payload.Password = utils.HashPassword(payload.Password)

	result = <-c.productRepositoryCommand.UpdateOneProduct(ctx, payload)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed update user")
		result.Error = errObj
		return result
	}

	return result
}

func (c commandUsecase) DeleteProduct(ctx context.Context, productId string) utils.Result {
	var result utils.Result

	queryRes := <-c.productRepositoryQuery.FindOne(ctx, productId)
	if queryRes.Data == nil {
		errObj := errors.NotFound("Product not found")
		result.Error = errObj
		return result
	}

	result = <-c.productRepositoryCommand.DeleteOneProduct(ctx, productId)
	if result.Error != nil {
		errObj := errors.InternalServerError("Failed delete product")
		result.Error = errObj
		return result
	}

	return result
}
