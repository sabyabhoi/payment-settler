package repositories

import (
	"gorm.io/gorm"
	"sabyabhoi.com/payment-settler/src/domain/models"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) FindAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction

	result := r.db.Find(&transactions)

	return transactions, result.Error
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) (uint, error) {
	result := r.db.Create(transaction)

	return transaction.ID, result.Error
}

func (r *TransactionRepository) UpdateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	result := r.db.First(transaction, transaction.ID)
	if result.Error != nil {
		return transaction, result.Error
	}

	result = r.db.Save(transaction)
	return transaction, result.Error
}

func (r *TransactionRepository) DeleteTransaction(groupId uint) error {
	result := r.db.Delete(&models.Transaction{}, groupId)

	return result.Error
}
