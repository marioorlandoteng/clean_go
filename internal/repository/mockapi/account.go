package mockapi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"moneytransfer/domain"
)

const (
	userURL = "https://663eb5dbe3a7c3218a4b345f.mockapi.io/api/v1/users"
)

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{}
}

func (m *AccountRepository) GetByNumberAndName(ctx context.Context, accountNumber string, accountName string) (domain.Account, error) {
	res, err := http.Get(userURL + "/" + accountNumber)
	if err != nil {
		return domain.Account{}, err 
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return domain.Account{}, err 
	}

    var account domain.Account
    err = json.Unmarshal([]byte(bodyBytes), &account)
    if err != nil {
		return domain.Account{}, err 
	}
	if (accountName != account.AccountName) {
		return domain.Account{}, errors.New("account not found")
	}
	return account, nil
}