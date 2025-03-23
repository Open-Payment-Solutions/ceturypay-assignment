package starters

import (
	"centurypay/internal/processors"
	"context"

	"centurypay/internal/di"
	"centurypay/internal/services"
)

func RegisterServices(_ctx context.Context, di di.Container) {
	accountsService := services.NewAccountsService()
	transactionsService := services.NewTransactionsService(accountsService)
	transactionsProcessor := processors.NewTransactionsProcessor(transactionsService)
	transactionsService.SetProcessor(transactionsProcessor)

	di.MustSet("accountsService", accountsService)
	di.MustSet("transactionsService", transactionsService)
	di.MustSet("transactionsProcessor", transactionsProcessor)

	SeedData(accountsService)
}
