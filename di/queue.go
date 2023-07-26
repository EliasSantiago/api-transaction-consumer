package di

import (
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/api-contatos/core/domain"
	"github.com/EliasSantiago/api-contatos/core/domain/usecase/transactionusecase"
)

func ConfigTransactionQueueDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionQueueRepository := transactionrepository.New(conn)
	transactionQueueUseCase := transactionusecase.New(transactionQueueRepository)
	return transactionQueueUseCase
}
