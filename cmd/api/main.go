package main

import (
	"context"

	"centurypay/internal/apps/api"
	"centurypay/internal/di"
	"centurypay/internal/handlers"
)

var (
	diContainer di.Container

	app *api.App
)

func main() {

	ctx := context.Background()

	diContainer = di.NewDI(ctx)
	app = api.NewApp(ctx, diContainer)

	registerDependencies(diContainer)

	handlers.StartLifecycle(app)
}

func registerDependencies(di di.Container) {
	di.MustSet("app", app)
}
