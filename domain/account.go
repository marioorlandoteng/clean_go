package domain

// Account represents bank account data struct
type Account struct {
	AccountNumber 	string `json:"accountNumber"`
	AccountName     string `json:"accountName"`
	CreatedAt 		string `json:"createdAt"`
}