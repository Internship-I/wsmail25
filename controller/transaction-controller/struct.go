package transaction_controller

import (
	"wsmail25/controller"
	"wsmail25/repository"
)

type TransactionHandler struct {
	transaction repository.TransactionRepository
}

func NewTransController(transaction repository.TransactionRepository) controller.TransactionController {
	return &TransactionHandler{
		transaction: transaction,
	}
}
