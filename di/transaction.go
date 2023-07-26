package di

import (
	"github.com/EliasSantiago/api-contatos/adapter/postgres"
	"github.com/EliasSantiago/api-contatos/adapter/postgres/transactionrepository"
	"github.com/EliasSantiago/api-contatos/core/domain"
	"github.com/EliasSantiago/api-contatos/core/domain/usecase/transactionusecase"
)

func ConfigTransactionDI(conn postgres.PoolInterface) domain.TransactionService {
	transactionRepository := transactionrepository.New(conn)
	transactionUseCase := transactionusecase.New(transactionRepository)
	return transactionUseCase
}
