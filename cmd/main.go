package main

import (
	"api_rest/cmd/docs"
	"api_rest/cmd/middlewares"
	"api_rest/cmd/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	router := gin.New()
	router.Use(middlewares.Logger())
	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routerInit := routes.NewRouter(router)
	routerInit.SetRoutes()

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
