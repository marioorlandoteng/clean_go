package mockapi

import (
	"context"

	"moneytransfer/domain"
)

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (m *AccountRepository) GetByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error) {
	// todo
	return domain.Account{}, nil
}