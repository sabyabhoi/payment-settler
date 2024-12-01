package handlers

import "sabyabhoi.com/payment-settler/src/domain/services"

type TransactionHandler struct {
	transactionService *services.TransactionService
}

func NewTransactionHandler(service *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: service}
}
