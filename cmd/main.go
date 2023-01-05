package main

import (
	"api_rest/cmd/handlers"
	errorpkg "api_rest/pkg/error"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var (
	errOpen   = errorpkg.CustomError{Msg: "Error: no se puede abrir el archivo"}
	errParser = errorpkg.CustomError{Msg: "Error: no se puede parsear el objeto"}
)

func main() {
	readFile()

	router := gin.Default()
	router.GET("/ping", handlers.Ping)
	routerProduct := router.Group("/products")
	routerProduct.GET("", handlers.GetProducts)
	routerProduct.GET("/:id", handlers.GetProduct)
	routerProduct.GET("/search", handlers.SearchProduct)
	routerProduct.POST("/", handlers.AddProduct)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func readFile() {
	data, err := ioutil.ReadFile("products.json")
	if err != nil {
		panic(errOpen)
	}

	err = json.Unmarshal(data, &handlers.ListProducts)
	if err != nil {
		panic(errParser)
	}
}
