package domain

import (
	"github.com/EliasSantiago/api-contatos/core/dto"
)

type Transaction struct {
	Id    int64   `json:"id"`
	Payer int64   `json:"payer"`
	Payee int64   `json:"payee"`
	Value float64 `json:"value"`
}

type TransactionService interface {
	Consumer()
}

type TransactionUseCase interface {
	Consumer()
}

type TransactionRepository interface {
	GetBalance(userID int64) (float64, error)
	UpdateBalance(transactionRequest *dto.UpdateBalanceStore, userID int64) error
	Update(transactionRequest *dto.UpdateTransactionStore) error
}
