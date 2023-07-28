package transactionrepository

import (
	"context"
	"log"

	"github.com/EliasSantiago/api-transaction-consumer/core/dto"
)

func (repository repository) Update(transactionRequest *dto.UpdateTransactionStore) error {
	ctx := context.Background()
	query := "UPDATE transactions SET status = $1 WHERE id = $2"
	_, err := repository.db.Exec(ctx, query, true, transactionRequest.ID)
	if err != nil {
		log.Printf("Erro ao executar a atualização no banco de dados: %v", err)
		return err
	}
	return nil
}
