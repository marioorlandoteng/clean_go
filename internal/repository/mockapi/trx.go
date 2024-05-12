package mockapi

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"moneytransfer/domain"
	"moneytransfer/internal/repository/postgresql"
)

const (
	bankURL = "https://663eb5dbe3a7c3218a4b345f.mockapi.io/api/v1/banks"
)

type TransactionRepository struct {
	httpClient	*http.Client
	dbRepo 		*postgresql.TransactionRepository 	// db proxy
}

func NewTransactionRepository(httpClient *http.Client, dbRepo *postgresql.TransactionRepository) *TransactionRepository {
	return &TransactionRepository{
		httpClient: httpClient,
		dbRepo: dbRepo,
	}
}

func (r *TransactionRepository) GetByRefID(ctx context.Context, refID string) (transaction domain.Transaction, err error) {
	transaction, err = r.dbRepo.GetByRefID(ctx, refID)
	if err != nil {
   		return domain.Transaction{}, err
   	}
	return transaction, nil
}

func (r *TransactionRepository) Create(ctx context.Context, transaction *domain.Transaction) error {
	// uuid as ref id
	transaction.RefID = uuid.New().String()

	var c chan error = make(chan error)

	go func(c chan error) {
		transactionJSON, err := json.Marshal(transaction)
		if err != nil {
			c <- err
			return
		}

		_, err = http.Post(bankURL, "application/json", bytes.NewBuffer(transactionJSON))
		if err != nil {
			c <- err
			return
		}
		c <- nil
	}(c)

	go func(c chan error) {
		err := r.dbRepo.Create(ctx, transaction)
		if err != nil {
	   		c <- err
			return
	   	}
		c <- nil
	}(c)
	
	for range 2 {
	    err := <- c
	    if (err != nil) {
	    	return err
	    }
	}

	return nil
}



func (r *TransactionRepository) UpdateStatus(ctx context.Context, transaction *domain.Transaction) error {
	err := r.dbRepo.UpdateStatus(ctx, transaction)
	if err != nil {
   		return err
   	}
	return nil
}