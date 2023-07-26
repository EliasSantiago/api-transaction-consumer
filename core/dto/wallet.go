package dto

import (
	"encoding/json"
	"io"
)

type UpdateWalletRequest struct {
	ID      int64   `json:"id" validate:"required,id"`
	UserID  int64   `json:"user_id" validate:"required,user_id"`
	Balance float64 `json:"balance" validate:"required,balance"`
}

type UpdateWalletStore struct {
	ID      int64   `db:"id"`
	UserID  int64   `db:"user_id"`
	Balance float64 `db:"balance"`
}

func FromJSONUpdateWalletRequest(body io.Reader) (*UpdateWalletRequest, error) {
	updateWalletRequest := UpdateWalletRequest{}
	if err := json.NewDecoder(body).Decode(&updateWalletRequest); err != nil {
		return nil, err
	}
	return &updateWalletRequest, nil
}
