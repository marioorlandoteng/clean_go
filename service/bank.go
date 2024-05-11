package service

import (
	"context"
	"moneytransfer/domain"
)

type AccountRepository interface {
	GetByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error)
}

type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
	Update(ctx context.Context, transaction *domain.Transaction) error
}

type BankService struct {
	accountRepo AccountRepository
	transactionRepo TransactionRepository
}

func NewBankService(accountRepo AccountRepository, transactionRepo TransactionRepository) *BankService {
	return &BankService{
		accountRepo: accountRepo,
		transactionRepo:  transactionRepo,
	}
}

func (b *BankService) GetAccountByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error) {
	// todo
	return domain.Account{}, nil
}  

func (b *BankService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}  

func (b *BankService) UpdateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	// todo
	return nil
}  