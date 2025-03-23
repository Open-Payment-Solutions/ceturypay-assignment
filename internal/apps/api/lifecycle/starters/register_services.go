package starters

import (
	"context"

	"centurypay/internal/di"
	"centurypay/internal/services"
)

func RegisterServices(_ctx context.Context, di di.Container) {
	di.MustSet("accountsService", services.NewAccountsService())
	di.MustSet("transactionsService", services.NewTransactionsService())
}
