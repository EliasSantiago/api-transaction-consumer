package transactionusecase

func (usecase usecase) GetBalance(userID int64) float64 {
	balance, err := usecase.repository.GetBalance(userID)
	if err != nil {
		return 0
	}
	return balance
}
