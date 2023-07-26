package transactionrepository

import (
	"context"
)

func (repository repository) GetBalance(userID int64) (float64, error) {
	ctx := context.Background()
	query := "SELECT balance FROM wallets WHERE user_id = $1"
	row := repository.db.QueryRow(ctx, query, userID)

	var balance float64
	err := row.Scan(&balance)
	if err != nil {
		return 0, nil
	}

	return balance, nil
}
