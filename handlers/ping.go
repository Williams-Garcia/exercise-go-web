package handlers

import (
	"api_rest/models"
	"api_rest/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var (
	nilList   = models.CustomError{Msg: "Lista nula"}
	errParser = models.CustomError{Msg: "Error al parsear"}
	errFound  = models.CustomError{Msg: "No se ha encontrado el objeto"}
)

var ListProducts = make([]models.Product, 0)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
	return
}

func GetProducts(ctx *gin.Context) {
	if ListProducts == nil {
		ctx.JSON(http.StatusConflict, nilList)
		return
	}
	ctx.JSON(http.StatusOK, models.Response{Data: ListProducts, Msg: "SUCCESS"})
	return
}

func GetProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errParser)
		return
	}

	if id == 0 {
		ctx.JSON(http.StatusNotFound, errFound)
		return
	}

	ctx.JSON(http.StatusFound, models.Response{Data: services.GetProductService(id, ListProducts), Msg: "SUCCESS"})
	return
}

func SearchProduct(ctx *gin.Context) {
	query, err := strconv.ParseFloat(ctx.Query("price"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errParser)
		return
	}
	ctx.JSON(http.StatusFound, models.Response{Data: services.SearchProductService(query, ListProducts), Msg: "SUCCESS"})
	return
}

func AddProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	product, err := services.AddProduct(product, ListProducts)
	if err != nil {
		if errors.Is(err, services.ErrItemExist) {
			ctx.JSON(http.StatusConflict, err)
			return
		}
		if errors.Is(err, services.ErrCodeValueRepeat) {
			ctx.JSON(http.StatusConflict, err)
			return
		}
		if errors.Is(err, services.ErrDateExp) {
			ctx.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ListProducts = append(ListProducts, product)

	ctx.JSON(http.StatusCreated, models.Response{Data: product, Msg: "SUCCESS"})
}
