package transactions

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
)

func (transactionRepo *Repository) GetAllTransactions() (*[]e.Transaction, error) {
	var transactions []e.Transaction
	if err := transactionRepo.DB.Preload("User", "deleted_at IS NULL").Preload("Payment", "deleted_at IS NULL").Preload("Transaction_details", "deleted_at IS NULL").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (transactionRepo *Repository) GetTransactionById(id int) (*e.Transaction, error) {
	var transaction e.Transaction
	if err := transactionRepo.DB.Preload("User", "deleted_at IS NULL").Preload("Payment", "deleted_at IS NULL").Preload("Transaction_details", "deleted_at IS NULL").Find(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transactionRepo *Repository) CreateTransaction(transaction *e.Transaction) error {
	if err := transactionRepo.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (transactionRepo *Repository) UpdateTransaction(id int, transaction *e.Transaction) error {
	result := transactionRepo.DB.Model(&transaction).Where("id = ?", id).Omit("UpdatedAt").Updates(&transaction)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (transactionRepo *Repository) DeleteTransaction(id int) error {
	if err := transactionRepo.DB.Where("transaction_id = ?", id).Delete(&e.TransactionDetail{}).Error; err != nil {
		return err
	}

	if err := transactionRepo.DB.Delete(&e.Transaction{}, id).Error; err != nil {
		return err
	}

	return nil
}
