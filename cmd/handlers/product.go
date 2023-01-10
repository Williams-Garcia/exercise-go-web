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

func NewProductHandler(ps impl.ProductService) *ProductHandler {
	return &ProductHandler{ProductService: ps}
}

func (ph *ProductHandler) GetProducts() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		// request

		// process
		products, err := ph.ProductService.GetProducts()
		if err != nil {
			ctx.JSON(http.StatusConflict, response.Err(err))
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok(products, "SUCCESS"))
	}
}

func (ph *ProductHandler) GetProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//	request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		//	process
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
}

func (ph *ProductHandler) AddProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//	request
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

		//process
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

		//	response
		ctx.JSON(http.StatusCreated, response.Ok(product, "SUCCESS"))
		return
	}
}

func (ph *ProductHandler) UpdateProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//	request
		var product domain.Product
		id, err := strconv.Atoi(ctx.Param("id"))

		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// validate := validator.New()
		// if err := validate.Struct(&product); err != nil {
		// 	ctx.JSON(http.StatusUnprocessableEntity, err)
		// 	return
		// }

		//process
		updatedProduct, err := ph.ProductService.UpdateProduct(id, product)
		if err != nil {
			if errors.Is(err, msg.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, response.Err(err))
				return
			}
			if errors.Is(err, msg.ErrCodeValueRepeat) {
				ctx.JSON(http.StatusConflict, response.Err(err))
				return
			}
			if errors.Is(err, msg.ErrDateExp) {
				ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		//	response
		ctx.JSON(http.StatusOK, response.Ok(updatedProduct, "SUCCESS"))
		return
	}
}

func (ph *ProductHandler) UpdatePatchProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//	request
		var productPatch domain.ProductPatch
		id, err := strconv.Atoi(ctx.Param("id"))

		if err := ctx.ShouldBind(&productPatch); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		// validate := validator.New()
		// if err := validate.Struct(&product); err != nil {
		// 	ctx.JSON(http.StatusUnprocessableEntity, err)
		// 	return
		// }

		//process
		updatedProduct, err := ph.ProductService.UpdatePatchProduct(id, productPatch)
		if err != nil {
			if errors.Is(err, msg.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, response.Err(err))
				return
			}
			if errors.Is(err, msg.ErrCodeValueRepeat) {
				ctx.JSON(http.StatusConflict, response.Err(err))
				return
			}
			if errors.Is(err, msg.ErrDateExp) {
				ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		//	response
		ctx.JSON(http.StatusOK, response.Ok(updatedProduct, "SUCCESS"))
		return
	}
}

func (ph *ProductHandler) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//	request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		//process
		err = ph.ProductService.DeleteProduct(id)

		if err != nil {
			if errors.Is(err, msg.ErrNotFound) {
				ctx.JSON(http.StatusNotFound, response.Err(err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		//	response
		ctx.JSON(http.StatusNoContent, response.Ok(nil, "SUCCESS"))
		return
	}
}

func (ph *ProductHandler) SearchProduct() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		//	request
		query, err := strconv.ParseFloat(ctx.Query("price"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, msg.ErrParser)
			return
		}

		//	process
		products, err := ph.ProductService.SearchProduct(query)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		//	response
		ctx.JSON(http.StatusOK, response.Ok(products, "SUCCESS"))
		return
	}
}
