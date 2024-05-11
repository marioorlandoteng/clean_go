package domain

// Transaction represents transfer/disbursement process data struct
type Transaction struct {
	ID        		int64 `json:"id"`
	FromUserID     	int64 `json:"fromUserID"`
	ToUserID 		int64 `json:"toUserID"`
	Amount			int64 `json:"amount"`
	Status			TransactionStatus `json:"transactionStatus"`
	CreatedAt 		string `json:"createdAt"`
	UpdatedAt 		string `json:"updatedAt"`
}
