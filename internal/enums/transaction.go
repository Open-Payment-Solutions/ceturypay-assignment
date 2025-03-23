package enums

type TransactionStatus string

const (
	TransactionStatusCreated   TransactionStatus = "created"
	TransactionStatusConfirmed TransactionStatus = "confirmed"
	TransactionStatusRejected  TransactionStatus = "rejected"
	TransactionStatusExpired   TransactionStatus = "expired"
)
