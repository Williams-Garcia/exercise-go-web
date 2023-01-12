package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	token := os.Getenv("TOKEN")

	return func(ctx *gin.Context) {
		requestToken := ctx.GetHeader("Token")

		if token != requestToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "NO AUTORIZADO")
			return
		}
		ctx.Next()
	}
}
