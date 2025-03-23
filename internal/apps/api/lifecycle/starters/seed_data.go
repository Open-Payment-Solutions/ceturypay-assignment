package starters

import (
	"centurypay/internal/enums"
	"centurypay/internal/interfaces"
)

func SeedData(accountsService interfaces.AccountsService) {
	_, _ = accountsService.CreateAccount("Mark", 100.0, enums.USD)
	_, _ = accountsService.CreateAccount("Jane", 50.0, enums.USD)
	_, _ = accountsService.CreateAccount("Adam", 0.0, enums.USD)
}
