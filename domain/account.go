package domain

// Account represents bank account data struct
type Account struct {
	ID        		int64  `json:"id"`
	AccountName     string `json:"accountName"`
	AccountNumber 	int64 `json:"accountNumber"`
	CreatedAt 		string `json:"createdAt"`
}