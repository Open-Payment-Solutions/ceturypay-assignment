package enums

type TransactionStatus string

const (
	TransactionStatusCreated   TransactionStatus = "created"
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusConfirmed TransactionStatus = "confirmed"
	TransactionStatusRejected  TransactionStatus = "rejected"
	TransactionStatusExpired   TransactionStatus = "expired"
	TransactionStatusCompleted TransactionStatus = "completed"
)
