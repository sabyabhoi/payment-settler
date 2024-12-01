package services

import (
	"sabyabhoi.com/payment-settler/src/domain/models"
	"sabyabhoi.com/payment-settler/src/infrastructure/repositories"
)

type TransactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: *repo}
}

func (s *TransactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.repo.FindAllTransactions()
}
