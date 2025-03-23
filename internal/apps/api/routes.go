package api

import (
	"centurypay/internal/services"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"centurypay/internal/di"
)

func RegisterRoutes(ctx context.Context, di di.Container, engine *echo.Echo) {
	accountsService := di.MustGet("accountsService").(*services.AccountsService)
	//transfersService := di.MustGet("transfersService").(*services.TransfersService)

	engine.GET("/accounts", func(c echo.Context) error {
		accounts := accountsService.GetAccounts()

		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{
				"items": accounts,
			},
			"  ",
		)
	})

	engine.GET("/accounts/:id", func(c echo.Context) error {
		account, err := accountsService.GetAccount(c.Param("id"))
		if err != nil {
			status := http.StatusBadRequest
			if err.Error() == services.ErrAccountNotFound.Error() {
				status = http.StatusNotFound
			}

			return c.JSONPretty(
				status,
				map[string]interface{}{
					"message": err.Error(),
				},
				"  ",
			)
		}

		return c.JSONPretty(
			http.StatusOK,
			account,
			"  ",
		)
	})

	engine.POST("/transfers", func(c echo.Context) error {
		//transfer := transfersService.CreateTransfer()

		// ToDO: Implement transfer creation
		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{},
			"  ",
		)
	})

	engine.DELETE("/transfers/:id", func(c echo.Context) error {
		//transfer := transfersService.GetTransfer(c.Param("id"))

		// ToDO: Implement transfer deletion
		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{},
			"  ",
		)
	})

	engine.GET("/transfers/:id", func(c echo.Context) error {
		//err := transfersService.DeleteTransfer(c.Param("id")) // delete only unconfirmed transfer

		// ToDO: Implement transfer retrieval
		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{},
			"  ",
		)
	})

	engine.POST("/transfers/:id/confirm", func(c echo.Context) error {
		//transfer, err := transfersService.ConfirmTransfer(c.Param("id"))

		// ToDO: Implement transfer confirmation
		return c.JSONPretty(
			http.StatusOK,
			map[string]interface{}{},
			"  ",
		)
	})
}
