package transactions

import (
	et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"
)

func (transactionDetailUsecase *Usecase) GetAllTransactionDetails() (*[]et.TransactionDetail, error) {
	transactionDetails, err := transactionDetailUsecase.Repository.GetAllTransactionDetails()
	return transactionDetails, err
}

func (transactionDetailUsecase *Usecase) GetTransactionDetailById(id int) (*et.TransactionDetail, error) {
	transactionDetail, err := transactionDetailUsecase.Repository.GetTransactionDetailById(id)
	return transactionDetail, err
}

func (transactionDetailUsecase *Usecase) CreateTransactionDetail(transactionDetail *et.TransactionDetail) error {
	err := transactionDetailUsecase.Repository.CreateTransactionDetail(transactionDetail)
	return err
}

func (transactionDetailUsecase *Usecase) UpdateTransactionDetail(id int, transactionDetail *et.TransactionDetail) error {
	result := transactionDetailUsecase.Repository.UpdateTransactionDetail(id, transactionDetail)
	return result
}

func (transactionDetailUsecase *Usecase) DeleteTransactionDetail(id int) error {
	err := transactionDetailUsecase.Repository.DeleteTransactionDetail(id)
	return err
}
