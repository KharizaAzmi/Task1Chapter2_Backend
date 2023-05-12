package handlers

import (
	"codebase-go/bin/middlewares"
	"codebase-go/bin/modules/product"
	"codebase-go/bin/modules/product/models"
	"codebase-go/bin/pkg/errors"
	"codebase-go/bin/pkg/helpers"

	"github.com/labstack/echo"
)

type productHttpHandler struct {
	productUsecaseQuery   product.UsecaseQuery
	productUseCaseCommand product.UsecaseCommand
}

func InitproductHttpHandler(e *echo.Echo, uq product.UsecaseQuery, uc product.UsecaseCommand) {

	handler := &productHttpHandler{
		productUsecaseQuery:   uq,
		productUseCaseCommand: uc,
	}

	route := e.Group("/codebase-go")

	route.GET("/product/v1", handler.GetListProduct, middlewares.VerifyBearer)
	route.POST("/product/v1/add", handler.CreateProduct, middlewares.VerifyBearer)
	route.PUT("/product/v1/update/:id", handler.UpdateProduct, middlewares.VerifyBearer)
	route.DELETE("/product/v1/delete/:id", handler.DeleteProduct, middlewares.VerifyBearer)
}

func (u productHttpHandler) GetListProduct(c echo.Context) error {
	productId := c.Get("productId").(string)
	result := u.productUsecaseQuery.GetDetailProduct(c.Request().Context(), productId)

	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Get list product")
}

func (u productHttpHandler) GetDetailProduct(c echo.Context) error {
	productId := c.Get("productId").(string)
	result := u.productUsecaseQuery.GetDetailProduct(c.Request().Context(), productId)

	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Get detail product")
}

func (u productHttpHandler) CreateProduct(c echo.Context) error {
	req := new(models.Product)

	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	// if err := c.Validate(req); err != nil {
	// 	return helpers.RespError(c, err)
	// }

	result := u.productUseCaseCommand.CreateProduct(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Add product success")
}

func (u productHttpHandler) UpdateProduct(c echo.Context) error {
	req := new(models.Product)
	req.Id = c.Param("id")

	if err := c.Bind(req); err != nil {
		return helpers.RespError(c, errors.BadRequest("bad request."))
	}
	if err := c.Validate(req); err != nil {
		return helpers.RespError(c, err)
	}
	if req.Id == "" {
		return helpers.RespError(c, errors.BadRequest("id can't empty."))
	}

	result := u.productUseCaseCommand.UpdateProduct(c.Request().Context(), *req)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Update Product success")
}

func (u productHttpHandler) DeleteProduct(c echo.Context) error {
	productId := c.Param("id")
	if productId == "" {
		return helpers.RespError(c, errors.BadRequest("id can't empty."))
	}

	result := u.productUseCaseCommand.DeleteProduct(c.Request().Context(), productId)
	if result.Error != nil {
		return helpers.RespError(c, result.Error)
	}

	return helpers.RespSuccess(c, result.Data, "Delete product success")
}
