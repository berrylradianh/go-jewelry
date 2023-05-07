package transactions

import (
	"fmt"

	et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"
)

func (transactionDetailRepo *Repository) GetAllTransactionDetails() (*[]et.TransactionDetail, error) {
	var transactionDetails []et.TransactionDetail
	if err := transactionDetailRepo.DB.Preload("Transaction", "deleted_at IS NULL").Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return &transactionDetails, nil
}

func (transactionDetailRepo *Repository) GetTransactionDetailById(id int) (*et.TransactionDetail, error) {
	var transactionDetail et.TransactionDetail
	if err := transactionDetailRepo.DB.Preload("Transaction", "deleted_at IS NULL").First(&transactionDetail, id).Error; err != nil {
		return nil, err
	}

	return &transactionDetail, nil
}

func (transactionDetailRepo *Repository) CreateTransactionDetail(transactionDetail *et.TransactionDetail) error {
	if err := transactionDetailRepo.DB.Create(&transactionDetail).Error; err != nil {
		return err
	}

	return nil
}

func (transactionDetailRepo *Repository) UpdateTransactionDetail(id int, transactionDetail *et.TransactionDetail) error {
	result := transactionDetailRepo.DB.Model(&transactionDetail).Where("id = ?", id).Omit("UpdatedAt").Updates(&transactionDetail)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (transactionDetailRepo *Repository) DeleteTransactionDetail(id int) error {
	if err := transactionDetailRepo.DB.Delete(&et.TransactionDetail{}, id).Error; err != nil {
		return err
	}

	return nil
}
