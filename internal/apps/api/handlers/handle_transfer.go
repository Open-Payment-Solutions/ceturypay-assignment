package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"centurypay/internal/interfaces"
)

type TransferRequest struct {
	FromAccountID string  `json:"fromAccountId"`
	ToAccountID   string  `json:"toAccountId"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

func HandleTransfer(c echo.Context, accountsService interfaces.AccountsService, transactionsService interfaces.TransactionsService) error {
	var req TransferRequest

	if err := c.Bind(&req); err != nil {
		return c.JSONPretty(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "Invalid request body",
			},
			"  ",
		)
	}

	if req.FromAccountID == "" || req.ToAccountID == "" || req.Amount <= 0 {
		return c.JSONPretty(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "Missing required fields or invalid amount",
			},
			"  ",
		)
	}

	fromAccount, err := accountsService.GetAccount(req.FromAccountID)
	if err != nil {
		return c.JSONPretty(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "From account not found",
			},
			"  ",
		)
	}

	toAccount, err := accountsService.GetAccount(req.ToAccountID)
	if err != nil {
		return c.JSONPretty(
			http.StatusBadRequest,
			map[string]interface{}{
				"message": "To account not found",
			},
			"  ",
		)
	}

	// Create transaction
	transaction, err := transactionsService.CreateTransaction(
		fromAccount,
		toAccount,
		req.Amount,
		req.Currency,
	)

	if err != nil {
		statusCode := http.StatusBadRequest
		message := err.Error()

		return c.JSONPretty(
			statusCode,
			map[string]interface{}{
				"message": message,
			},
			"  ",
		)
	}

	return c.JSONPretty(
		http.StatusCreated,
		transaction,
		"  ",
	)
}
