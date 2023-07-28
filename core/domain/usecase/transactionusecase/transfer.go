package transactionusecase

import (
	"github.com/EliasSantiago/api-transaction-consumer/core/dto"
)

func (usecase usecase) UpdateBalance(transactionRequest *dto.UpdateTransactionRequest, userID int64) error {
	store := &dto.UpdateBalanceStore{
		ID:      transactionRequest.ID,
		Balance: transactionRequest.Value,
	}
	err := usecase.repository.UpdateBalance(store, userID)
	if err != nil {
		return err
	}
	return nil
}
