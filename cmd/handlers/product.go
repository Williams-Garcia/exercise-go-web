package handlers

import (
	"api_rest/internal/domain"
	msg "api_rest/internal/product"
	"api_rest/internal/product/impl"
	"api_rest/pkg/response"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ProductHandler struct {
	ProductService impl.ProductService
}

func (ph *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := ph.ProductService.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusConflict, response.Err(err))
	}

	ctx.JSON(http.StatusOK, response.Ok(products, "SUCCESS"))
}

func (ph *ProductHandler) GetProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Err(err))
		return
	}

	product, err := ph.ProductService.GetProduct(id)

	if err != nil {
		if errors.Is(err, msg.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, response.Err(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusFound, response.Ok(product, "SUCCESS"))
	return
}

func (ph *ProductHandler) AddProduct(ctx *gin.Context) {
	var product domain.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	product, err := ph.ProductService.AddProduct(product)
	if err != nil {
		if errors.Is(err, msg.ErrItemExist) {
			ctx.JSON(http.StatusConflict, err)
			return
		}
		if errors.Is(err, msg.ErrCodeValueRepeat) {
			ctx.JSON(http.StatusConflict, err)
			return
		}
		if errors.Is(err, msg.ErrDateExp) {
			ctx.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, response.Ok(product, "SUCCESS"))
	return
}

func (ph *ProductHandler) SearchProduct(ctx *gin.Context) {
	query, err := strconv.ParseFloat(ctx.Query("price"), 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, msg.ErrParser)
		return
	}

	products, err := ph.ProductService.SearchProduct(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Err(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Ok(products, "SUCCESS"))
	return
}
