package transaction_controller

import (
	// "wsmail25/c"
	"wsmail25/repository"
)

type TransactionHandler struct {
	transaction repository.TransactionRepository
}

func NewTransController(transaction repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{
		transaction: transaction,
	}
}
