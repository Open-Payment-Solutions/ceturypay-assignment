package api

import (
	"centurypay/internal/apps/api/handlers"
	"centurypay/internal/interfaces"
	"centurypay/internal/services"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"centurypay/internal/di"
)

func RegisterRoutes(ctx context.Context, di di.Container, engine *echo.Echo) {
	accountsService := di.MustGet("accountsService").(interfaces.AccountsService)
	transactionsService := di.MustGet("transactionsService").(interfaces.TransactionsService)

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
		return handlers.HandleTransfer(c, accountsService, transactionsService)
	})

	// Add endpoint to get transaction status
	engine.GET("/transactions/:id", func(c echo.Context) error {
		transaction, err := transactionsService.GetTransaction(c.Param("id"))
		if err != nil {
			status := http.StatusBadRequest
			if err.Error() == services.ErrTransactionNotFound.Error() {
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
			transaction,
			"  ",
		)
	})
}
