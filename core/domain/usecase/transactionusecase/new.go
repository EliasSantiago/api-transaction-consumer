package transactionusecase

import "github.com/EliasSantiago/api-contatos/core/domain"

type usecase struct {
	repository domain.TransactionRepository
}

func New(repository domain.TransactionRepository) domain.TransactionUseCase {
	return &usecase{
		repository: repository,
	}
}
