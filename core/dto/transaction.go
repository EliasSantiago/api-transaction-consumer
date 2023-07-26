package dto

import (
	"encoding/json"
	"io"
)

type UpdateTransactionRequest struct {
	ID    int64   `json:"id" validate:"required,id"`
	Payer int64   `json:"payer" validate:"required,payer"`
	Payee int64   `json:"payee" validate:"required,payee"`
	Value float64 `json:"value" validate:"required,value"`
}

type GetBalanceRequest struct {
	Payer int64 `json:"payer" validate:"required,payer"`
}

// type UpdateTransactionResponse struct {
// 	ID        string `json:"id"`
// 	Name      string `json:"name"`
// 	LastName  string `json:"last_name"`
// 	Phone     string `json:"phone"`
// 	Address   string `json:"address"`
// 	DateBirth string `json:"date_birth"`
// 	Cpf       string `json:"cpf"`
// }

type UpdateTransactionStore struct {
	ID    int64   `db:"id"`
	Payer int64   `db:"payer"`
	Payee int64   `db:"payee"`
	Value float64 `db:"value"`
}

type UpdateBalanceStore struct {
	ID      int64   `db:"id"`
	Balance float64 `db:"balance"`
}

func FromJSONUpdateTransactionRequest(body io.Reader) (*UpdateTransactionRequest, error) {
	updateTransactionRequest := UpdateTransactionRequest{}
	if err := json.NewDecoder(body).Decode(&updateTransactionRequest); err != nil {
		return nil, err
	}
	return &updateTransactionRequest, nil
}
