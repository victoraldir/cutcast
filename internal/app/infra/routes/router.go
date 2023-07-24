package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victoraldir/cutcast/internal/app/infra/config"
)

func DefaultRouter(cfg config.Configuration) *gin.Engine {

	router := gin.New()
	gin.SetMode(gin.DebugMode)
	router.NoRoute(noRouteHandler)

	return router
}

func noRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": fmt.Sprintf("resource %s not found", ctx.Request.URL.Path),
	})
}
