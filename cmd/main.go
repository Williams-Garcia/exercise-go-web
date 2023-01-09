package main

import (
	"api_rest/cmd/handlers"
	"api_rest/internal/product"
	"api_rest/internal/product/impl"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/ping", handlers.Ping)
	routerProduct := router.Group("/products")

	repo := product.NewProductRepository()
	service := *impl.NewProductService(repo)
	serviceImpl := handlers.ProductHandler{ProductService: service}

	routerProduct.GET("", serviceImpl.GetProducts)
	routerProduct.GET("/:id", serviceImpl.GetProduct)
	routerProduct.GET("/search", serviceImpl.SearchProduct)
	routerProduct.POST("/", serviceImpl.AddProduct)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
