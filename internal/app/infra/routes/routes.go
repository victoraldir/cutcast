package routes

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/victoraldir/cutcast/internal/app"
)

func MapRoutes(router *gin.Engine, app *app.Application) {
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/numgoroutines", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"num_goroutines": runtime.NumGoroutine(),
		})
	})

	v1 := router.Group("/v1")
	{
		record := v1.Group("/record")
		{
			record.POST("", app.RecordGroupController.Create)
		}

		recordName := v1.Group("/record/:id")
		{
			recordName.PUT("/finish", app.RecordGroupController.Finish)

			trimGroup := recordName.Group("/trim")
			{
				trimGroup.POST("", app.TrimGroupController.Create)
				trimGroup.GET("", app.TrimGroupController.List)
			}
		}

		recordList := v1.Group("/record")
		{
			recordList.GET("", app.RecordGroupController.List)
		}
	}
}
