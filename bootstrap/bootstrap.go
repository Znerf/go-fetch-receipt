package bootstrap

import (
	"context"
	"receipt/api/controllers"
	"receipt/resource"
	"receipt/api/routes"
	"receipt/api/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	resource.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler resource.Router,
	routes routes.Routes,
) {
	appStop := func(context.Context) error {
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				routes.Setup()
				handler.Gin.Run(":" + "8000")
			}()
			return nil
		},
		OnStop: appStop,
	})
}