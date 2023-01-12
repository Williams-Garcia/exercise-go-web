package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				typeRequest := ctx.Request.Method
				timeRequest := time.Now()
				urlRequest := ctx.Request.URL
				sizeRequest := ctx.Request.ContentLength

				log.Printf("Panic: %s, en metodo %s a las %v. Endpoint: %s, Size: %d", err, typeRequest, timeRequest, urlRequest, sizeRequest)
			}
		}()

		ctx.Next()
	}
}
