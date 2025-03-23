package services

type TransactionsService struct {
}

func NewTransactionsService() *TransactionsService {
	return &TransactionsService{}
}

func (s *TransactionsService) Stop() error {
	// ToDo: Implement the logic to stop the service
	return nil
}
