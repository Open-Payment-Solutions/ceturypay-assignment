package api

import (
	"centurypay/internal/apps/api/lifecycle/starters"
	"centurypay/internal/apps/api/lifecycle/stoppers"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"

	"centurypay/internal/di"
)

type App struct {
	ctx    context.Context
	di     di.Container
	engine *echo.Echo
}

func NewApp(ctx context.Context, di di.Container) *App {
	engine := echo.New()
	engine.HideBanner = true

	starters.RegisterServices(ctx, di)

	RegisterMiddlewares(ctx, di, engine)
	RegisterRoutes(ctx, di, engine)

	return &App{
		ctx:    ctx,
		di:     di,
		engine: engine,
	}
}

func (app *App) SetListener(listener net.Listener) {
	app.engine.Listener = listener
}

func (app *App) SetServer(server *http.Server) {
	app.engine.Server = server
}

func (app *App) Start() error {
	listenAt := app.ctx.Value("ListenAt").(string)

	err := app.startListener(listenAt)
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (app *App) startListener(listenAt string) error {
	return app.engine.Start(listenAt)
}

func (app *App) Stop() error {
	stoppers.StopServices(app.ctx, app.di)

	return app.engine.Shutdown(app.ctx)
}
