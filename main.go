package main

import (
	"api_rest/handlers"
	"api_rest/models"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

var (
	errOpen   = models.CustomError{Msg: "Error: no se puede abrir el archivo"}
	errParser = models.CustomError{Msg: "Error: no se puede parsear el objeto"}
)

func main() {
	readFile()

	router := gin.Default()
	router.GET("/ping", handlers.Ping)
	routerProduct := router.Group("/products")
	routerProduct.GET("", handlers.GetProducts)
	routerProduct.GET("/:id", handlers.GetProduct)
	routerProduct.GET("/search", handlers.SearchProduct)

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
