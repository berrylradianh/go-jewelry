package transactions

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (transactionDetailUsecase *Usecase) GetAllTransactionDetails() (*[]e.TransactionDetail, error) {
	transactionDetails, err := transactionDetailUsecase.Repository.GetAllTransactionDetails()
	return transactionDetails, err
}

func (transactionDetailUsecase *Usecase) GetTransactionDetailById(id int) (*e.TransactionDetail, error) {
	transactionDetail, err := transactionDetailUsecase.Repository.GetTransactionDetailById(id)
	return transactionDetail, err
}

func (transactionDetailUsecase *Usecase) FindProduct(id int) (*e.Product, error) {
	product, err := transactionDetailUsecase.Repository.FindProduct(id)
	return product, err
}

func (transactionDetailUsecase *Usecase) CreateTransactionDetail(transactionDetail *e.TransactionDetail) error {
	err := transactionDetailUsecase.Repository.CreateTransactionDetail(transactionDetail)
	return err
}

func (transactionDetailUsecase *Usecase) UpdateTransactionDetail(id int, transactionDetail *e.TransactionDetail) error {
	result := transactionDetailUsecase.Repository.UpdateTransactionDetail(id, transactionDetail)
	return result
}

func (transactionDetailUsecase *Usecase) DeleteTransactionDetail(id int) error {
	err := transactionDetailUsecase.Repository.DeleteTransactionDetail(id)
	return err
}
