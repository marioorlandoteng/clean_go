package service

import (
	"context"

	"moneytransfer/domain"
)

type AccountRepository interface {
	GetByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error)
}

type TransactionRepository interface {
	GetByRefID(ctx context.Context, refID string) (transaction domain.Transaction, err error)
	Create(ctx context.Context, transaction *domain.Transaction) error
	UpdateStatus(ctx context.Context, transaction *domain.Transaction) error
}

type BankService struct {
	accountRepo AccountRepository
	transactionRepo TransactionRepository
}

func NewBankService(accountRepo AccountRepository, transactionRepo TransactionRepository) *BankService {
	return &BankService{
		accountRepo: accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *BankService) GetAccountByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error) {
	account, err := s.accountRepo.GetByNumberAndName(ctx, accountNumber, accountName)
	if err != nil {
		return domain.Account{}, err
   	}
	return account, nil
}  

func (s *BankService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	transaction.Status = domain.TRANSFERRED 
	err := s.transactionRepo.Create(ctx, transaction)
	if err != nil {
   		return err
   	}
	return nil
}  

func (s *BankService) DisburseTransaction(ctx context.Context, refID string) error {
	transaction, err := s.transactionRepo.GetByRefID(ctx, refID)
   	if err != nil {
   		return err
   	}

	transaction.Status = domain.DISBURSED
	err = s.transactionRepo.UpdateStatus(ctx, &transaction)
	if err != nil {
   		return err
   	}
	return nil
}  