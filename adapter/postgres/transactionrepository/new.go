package transactionrepository

import (
	"github.com/EliasSantiago/api-transaction-consumer/adapter/postgres"
	"github.com/EliasSantiago/api-transaction-consumer/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

func New(db postgres.PoolInterface) domain.TransactionRepository {
	return &repository{
		db: db,
	}
}
