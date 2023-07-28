package di

import (
	"github.com/EliasSantiago/api-transaction-consumer/adapter/postgres"
	"github.com/EliasSantiago/api-transaction-consumer/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/api-transaction-consumer/core/domain"
	"github.com/EliasSantiago/api-transaction-consumer/core/domain/usecase/transactionusecase"
)

func ConfigTransactionDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionRepository := transactionrepository.New(conn)
	transactionUseCase := transactionusecase.New(transactionRepository)
	return transactionUseCase
}
