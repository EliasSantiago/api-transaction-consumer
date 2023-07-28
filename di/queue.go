package di

import (
	"github.com/EliasSantiago/api-transaction-consumer/adapter/postgres"
	"github.com/EliasSantiago/api-transaction-consumer/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/api-transaction-consumer/core/domain"
	"github.com/EliasSantiago/api-transaction-consumer/core/domain/usecase/transactionusecase"
)

func ConfigTransactionQueueDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionQueueRepository := transactionrepository.New(conn)
	transactionQueueUseCase := transactionusecase.New(transactionQueueRepository)
	return transactionQueueUseCase
}
