package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
	return
}
