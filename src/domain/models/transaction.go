package models

import "time"

type Status string

const (
	STATUS_SETTLED Status = "SETTLED"
	STATUS_PENDING Status = "PENDING"
)

type Transaction struct {
	ID               uint
	Group            Group
	Payer            User
	Description      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Status           Status            `gorm:"not null"`
	TransactionInfos []TransactionInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type TransactionInfo struct {
	ID          uint
	Transaction Transaction
	Payee       User
	Amount      float32 `gorm:"check:amount > 0"`
	Status      Status
}
