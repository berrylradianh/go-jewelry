package transactions

import (
	"time"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (transactionUsecase *Usecase) GetAllTransactions() (*[]e.Transaction, error) {
	transactions, err := transactionUsecase.Repository.GetAllTransactions()
	return transactions, err
}

func (transactionUsecase *Usecase) GetTransactionById(id int) (*e.Transaction, error) {
	transaction, err := transactionUsecase.Repository.GetTransactionById(id)
	return transaction, err
}

func (transactionUsecase *Usecase) CreateTransaction(transaction *e.Transaction) error {
	transaction.Date = time.Now()
	transaction.Status = "Process"
	err := transactionUsecase.Repository.CreateTransaction(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (transactionUsecase *Usecase) UpdateTransaction(id int, transaction *e.Transaction) error {
	result := transactionUsecase.Repository.UpdateTransaction(id, transaction)
	return result
}

func (transactionUsecase *Usecase) DeleteTransaction(id int) error {
	err := transactionUsecase.Repository.DeleteTransaction(id)
	return err
}
