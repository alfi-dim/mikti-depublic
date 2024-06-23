package web

type TransactionRequest struct {
	EventID  string `json:"event_id"`
	Quantity int    `json:"quantity"`
}

type ConfirmPaymentRequest struct {
	TransactionId string `json:"transaction_id"`
}
