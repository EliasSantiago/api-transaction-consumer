package transactionusecase

import (
	"encoding/json"
	"fmt"

	"github.com/EliasSantiago/api-contatos/core/dto"
	"github.com/EliasSantiago/api-contatos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (usecase usecase) Consumer() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	numWorkers := 2
	for i := 1; i <= numWorkers; i++ {
		for msg := range out {
			inputDTO := &dto.UpdateTransactionRequest{}
			if err := json.Unmarshal(msg.Body, &inputDTO); err != nil {
				panic(err)
			}

			id := inputDTO.ID
			userPayer := inputDTO.Payer
			userPayee := inputDTO.Payee

			store := &dto.UpdateTransactionRequest{
				ID:    id,
				Payer: inputDTO.Payer,
				Payee: inputDTO.Payee,
				Value: inputDTO.Value,
			}

			balancePayer := usecase.GetBalance(userPayer)
			store.Value = balancePayer - inputDTO.Value
			err = usecase.UpdateBalance(store, userPayer)
			if err != nil {
				panic(err)
			}

			balance := usecase.GetBalance(userPayee)
			store.Value = balance + inputDTO.Value
			err = usecase.UpdateBalance(store, userPayee)
			if err != nil {
				panic(err)
			}

			err = usecase.Update(store)
			if err != nil {
				// TODO: TRATAR PARA PUBLICAR EM UMA FILA DE ERROS
				panic(err)
			}
			msg.Ack(false)
			fmt.Printf("Worker %d has processed transaction_id: %d\n", i, id)
		}
	}
}
