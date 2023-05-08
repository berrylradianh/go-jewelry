package transactions

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (transactionDetailRepo *Repository) GetAllTransactionDetails() (*[]e.TransactionDetail, error) {
	var transactionDetails []e.TransactionDetail
	if err := transactionDetailRepo.DB.Preload("Transaction", "deleted_at IS NULL").Preload("Product", "deleted_at IS NULL").Find(&transactionDetails).Error; err != nil {
		return nil, err
	}

	return &transactionDetails, nil
}

func (transactionDetailRepo *Repository) GetTransactionDetailById(id int) (*e.TransactionDetail, error) {
	var transactionDetail e.TransactionDetail
	if err := transactionDetailRepo.DB.Preload("Transaction", "deleted_at IS NULL").Preload("Product", "deleted_at IS NULL").First(&transactionDetail, id).Error; err != nil {
		return nil, err
	}

	return &transactionDetail, nil
}

func (transactionDetailRepo *Repository) FindProduct(id int) (*e.Product, error) {
	var product e.Product
	if err := transactionDetailRepo.DB.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (transactionDetailRepo *Repository) CreateTransactionDetail(transactionDetail *e.TransactionDetail) error {
	if err := transactionDetailRepo.DB.Create(&transactionDetail).Error; err != nil {
		return err
	}

	return nil
}

func (transactionDetailRepo *Repository) UpdateTransactionDetail(id int, transactionDetail *e.TransactionDetail) error {
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
	if err := transactionDetailRepo.DB.Delete(&e.TransactionDetail{}, id).Error; err != nil {
		return err
	}

	return nil
}
