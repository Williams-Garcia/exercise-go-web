package main

import (
	"api_rest/cmd/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routerInit := routes.NewRouter(router)
	routerInit.SetRoutes()

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
