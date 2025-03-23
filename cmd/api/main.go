package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"

	"centurypay/internal/apps/api"
	"centurypay/internal/di"
	"centurypay/internal/handlers"
)

var (
	diContainer di.Container

	app *api.App
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, "ListenAt", os.Getenv("LISTEN_AT"))

	diContainer = di.NewDI(ctx)
	app = api.NewApp(ctx, diContainer)

	registerDependencies(diContainer)

	handlers.StartLifecycle(app)
}

func registerDependencies(di di.Container) {
	di.MustSet("app", app)
}
