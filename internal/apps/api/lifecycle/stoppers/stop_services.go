package stoppers

import (
	"context"

	"centurypay/internal/di"
	"centurypay/internal/interfaces"
)

func StopServices(_ctx context.Context, di di.Container) {
	transactionsProcessor := di.MustGet("transactionsProcessor").(interfaces.TransactionProcessor)
	_ = transactionsProcessor.Stop()

	accountsService := di.MustGet("accountsService").(interfaces.AccountsService)
	_ = accountsService.Stop()

	transfersService := di.MustGet("transactionsService").(interfaces.TransactionService)
	_ = transfersService.Stop()
}
