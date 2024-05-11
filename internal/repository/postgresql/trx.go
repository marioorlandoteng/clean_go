package postgresql

import (
	"context"
	"database/sql"

	"moneytransfer/domain"
)

type TransactionRepository struct {
	Conn *sql.DB
}

func NewTransactionRepository(conn *sql.DB) *TransactionRepository {
	return &TransactionRepository{conn}
}

func (m *TransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}

func (m *TransactionRepository) Update(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}