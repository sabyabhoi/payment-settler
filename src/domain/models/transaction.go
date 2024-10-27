package models

import "time"

type Transaction struct {
	ID        uint
	Payer     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionInfo struct {
	ID            uint
	TransactionId uint
	PayeeId       uint
	Amount        float32
}
