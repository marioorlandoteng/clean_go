package domain

// Transaction represents transfer/disbursement process data struct
type Transaction struct {
	ID        			int64 `json:"id" db:"id"`
	FromAccountNo     	int64 `json:"fromAccountNo" db:"from_account_no" validate:"required"`
	ToAccountNo 		int64 `json:"toAccountNo" db:"to_account_no" validate:"required"`
	Amount				int `json:"amount" db:"amount" validate:"required"`
	Status				TransactionStatus `json:"status" db:"status"`
	CreatedAt 			int64 `json:"createdAt" db:"created_at"`
	UpdatedAt 			int64 `json:"updatedAt" db:"updated_at"`
	RefID        		string `json:"refId" db:"ref_id"`
}
