package product

import (
	"context"

	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/utils"
)

type UsecaseQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	GetDetailProduct(ctx context.Context, productId string) utils.Result
}

type UsecaseCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	CreateProduct(ctx context.Context, payload models.Product) utils.Result
	UpdateProduct(ctx context.Context, payload models.Product) utils.Result
	DeleteProduct(ctx context.Context, productId string) utils.Result
}

type MongodbRepositoryQuery interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	FindOne(ctx context.Context, userId string) <-chan utils.Result
	FindOneByUsername(ctx context.Context, username string) <-chan utils.Result
}

type MongodbRepositoryCommand interface {
	// idiomatic go, ctx first before payload. See https://pkg.go.dev/context#pkg-overview
	NewObjectID(ctx context.Context) string
	InsertOneProduct(ctx context.Context, data models.Product) <-chan utils.Result
	UpdateOneProduct(ctx context.Context, data models.Product) <-chan utils.Result
	DeleteOneProduct(ctx context.Context, productId string) <-chan utils.Result
}
