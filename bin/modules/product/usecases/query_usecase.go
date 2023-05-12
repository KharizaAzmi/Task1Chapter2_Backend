package usecases

import (
	"context"

	"codebase-go/bin/modules/product"
	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/redis"
	"codebase-go/bin/pkg/utils"
)

type queryUsecase struct {
	productRepositoryQuery product.MongodbRepositoryQuery
}

func NewQueryUsecase(mq product.MongodbRepositoryQuery, rc redis.Collections) product.UsecaseQuery {
	return &queryUsecase{
		productRepositoryQuery: mq,
	}
}

func (q queryUsecase) GetListProduct(ctx context.Context, productId string) utils.Result {
	var result utils.Result

	queryRes := <-q.productRepositoryQuery.FindOne(ctx, productId)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}
	product := queryRes.Data.(models.Product)
	res := models.GetProductResponse{
		Id: product.Id,                 
		Productname: product.ProductName,         
		Price: product.Price,               
		// OptionPaymentMethod: product.OptionPaymentMethod, 
		// DeliveryMethod: product.DeliveryMethod,      
		// StockDescription: product.StockDescription,    
		Image: product.Image,               
	}
	result.Data = res
	return result
}

func (q queryUsecase) GetDetailProduct(ctx context.Context, productId string) utils.Result {
	var result utils.Result

	queryRes := <-q.productRepositoryQuery.FindOne(ctx, productId)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}
	product := queryRes.Data.(models.Product)
	res := models.GetProductResponse{
		Id: product.Id,                 
		Productname: product.ProductName,         
		Price: product.Price,               
		OptionPaymentMethod: product.OptionPaymentMethod, 
		DeliveryMethod: product.DeliveryMethod,      
		StockDescription: product.StockDescription,    
		Image: product.Image,               
	}
	result.Data = res
	return result
}

func (q queryUsecase) FindOneByProductname(ctx context.Context, Productname string) utils.Result {
	var result utils.Result

	queryRes := <-q.productRepositoryQuery.FindOneByUsername(ctx, Productname)

	if queryRes.Error != nil {
		errObj := errors.InternalServerError("Internal server error")
		result.Error = errObj
		return result
	}
	product := queryRes.Data.(models.Product)
	res := models.GetProductResponse{
		Id: product.Id,                 
		Productname: product.ProductName,         
		Price: product.Price,               
		OptionPaymentMethod: product.OptionPaymentMethod, 
		DeliveryMethod: product.DeliveryMethod,      
		StockDescription: product.StockDescription,    
		Image: product.Image,               
	}
	result.Data = res
	return result
}
