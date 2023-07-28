package transactionrepository

import (
	"context"
	"log"

	"github.com/EliasSantiago/api-transaction-consumer/core/dto"
)

func (repository repository) UpdateBalance(walletRequest *dto.UpdateBalanceStore, userID int64) error {
	ctx := context.Background()
	query := "UPDATE wallets SET balance = $1 WHERE user_id = $2"
	_, err := repository.db.Exec(ctx, query, walletRequest.Balance, userID)
	if err != nil {
		log.Printf("Erro ao executar a atualização no banco de dados: %v", err)
		return err
	}
	return nil
}
