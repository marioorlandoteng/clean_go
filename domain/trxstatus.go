package domain

type TransactionStatus int

const (
	TRANSFERRED TransactionStatus = iota
	DISBURSED	
)