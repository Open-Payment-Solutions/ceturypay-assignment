package stoppers

import (
	"centurypay/internal/di"
	"centurypay/internal/services"
	"context"
)

func StopServices(_ctx context.Context, di di.Container) {
	accountsService := di.MustGet("accountsService").(*services.AccountsService)
	_ = accountsService.Stop()

	transfersService := di.MustGet("transfersService").(*services.TransfersService)
	_ = transfersService.Stop()
}
