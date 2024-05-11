package mockapi

import (
	"context"

	"moneytransfer/domain"
	"moneytransfer/internal/repository/postgresql"
)

type TransactionRepository struct {
	dbRepo *postgresql.TransactionRepository
}

func NewTransactionRepository(dbRepo *postgresql.TransactionRepository) *TransactionRepository {
	return &TransactionRepository{
		dbRepo: dbRepo,
	}
}

func (m *TransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}

func (m *TransactionRepository) Update(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}