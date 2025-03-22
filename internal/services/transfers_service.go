package services

type TransfersService struct {
}

func NewTransferService() *TransfersService {
	return &TransfersService{}
}

func (s *TransfersService) Stop() error {
	// ToDo: Implement the logic to stop the service
	return nil
}
