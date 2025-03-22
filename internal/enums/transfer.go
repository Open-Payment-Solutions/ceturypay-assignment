package enums

type TransferStatus = string

const (
	TransferPending TransferStatus = "PENDING"
	TransferSuccess TransferStatus = "SUCCESS"
	TransferFailed  TransferStatus = "FAILED"
)
