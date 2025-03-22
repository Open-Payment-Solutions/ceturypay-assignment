package api

import (
	"centurypay/internal/di"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func RegisterMiddlewares(ctx context.Context, di di.Container, engine *echo.Echo) {
	attachRequestIdMiddleware(engine)

	engine.Pre(middleware.RemoveTrailingSlash())
	engine.Use(middleware.Recover())

	engine.Use(
		middleware.LoggerWithConfig(
			middleware.LoggerConfig{
				Format: `${time_rfc3339} | ${remote_ip} | ${method} ${uri} | ${status} | ${latency_human} | ${bytes_in} / ${bytes_out} (in/out) | ${id}` + "\n",
			},
		),
	)

	postResponseLogger(engine)
}

func postResponseLogger(engine *echo.Echo) {
	engine.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lastLog := c.Get("LastLog")
			if lastLog != nil && strings.Trim(lastLog.(string), " ") != "" {
				fmt.Println(lastLog)
			}
			return next(c)
		}
	})
}

func attachRequestIdMiddleware(engine *echo.Echo) {
	config := middleware.RequestIDConfig{
		TargetHeader: "Request-Id",
		RequestIDHandler: func(c echo.Context, requestId string) {
			c.Set("RequestId", requestId)
			c.Request().Header.Set("X-Request-Id", requestId)
		},
	}
	engine.Pre(middleware.RequestIDWithConfig(config))
}
