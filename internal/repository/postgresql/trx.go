package postgresql

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"moneytransfer/domain"
)

const (
	tableName = "transaction"
)

type TransactionRepository struct {
	conn *sqlx.DB
}

func NewTransactionRepository(conn *sqlx.DB) *TransactionRepository {
	return &TransactionRepository{
		conn: conn,
	}
}

func (r *TransactionRepository) GetByRefID(ctx context.Context, refID string) (transaction domain.Transaction, err error) {
	query := `SELECT * FROM transaction WHERE ref_id = $1`

	err = r.conn.QueryRowx(query, refID).StructScan(&transaction)
   	if err != nil {
   		return domain.Transaction{}, err
   	}
	return transaction, nil
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	transaction.CreatedAt = time.Now().UnixMilli()
	transaction.UpdatedAt = time.Now().UnixMilli()

	query := `INSERT INTO transaction(from_account_no, to_account_no, amount, status, created_at, updated_at, ref_id) 
			  VALUES(:from_account_no, :to_account_no, :amount, :status, :created_at, :updated_at, :ref_id)`

	_, err := r.conn.NamedExec(query, transaction)
   	if err != nil {
   		return err
   	}
	return nil
}

func (r *TransactionRepository) UpdateStatus(ctx context.Context, transaction *domain.Transaction) error {
	transaction.UpdatedAt = time.Now().UnixMilli()

	query := `UPDATE transaction SET status = :status, updated_at = :updated_at WHERE id = :id`

	_, err := r.conn.NamedExec(query, transaction)
   	if err != nil {
   		return err
   	}
	return nil
}