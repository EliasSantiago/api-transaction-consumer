package transactionusecase

import (
	"github.com/EliasSantiago/api-contatos/core/dto"
)

func (usecase usecase) Update(transactionRequest *dto.UpdateTransactionRequest) error {
	store := &dto.UpdateTransactionStore{
		ID:    transactionRequest.ID,
		Payer: transactionRequest.Payer,
		Payee: transactionRequest.Payee,
		Value: transactionRequest.Value,
	}
	err := usecase.repository.Update(store)
	if err != nil {
		return err
	}
	return nil
}
