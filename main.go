package main

import (
	"github.com/victoraldir/cutcast/internal/app"
	"github.com/victoraldir/cutcast/internal/app/infra/config"
	"github.com/victoraldir/cutcast/internal/app/infra/routes"
)

func main() {
	cfg := config.InitConfiguration()
	app := app.NewApplication(cfg)

	router := routes.DefaultRouter(cfg)
	routes.MapRoutes(router, app)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
