package transactions

import (
	"fmt"

	et "github.com/berrylradianh/go-jewelry/modules/entity/transactions"
)

func (transactionRepo *Repository) GetAllTransactions() (*[]et.Transaction, error) {
	var transactions []et.Transaction
	if err := transactionRepo.DB.Preload("Transaction_detail", "deleted_at IS NULL").Find(&transactions).Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (transactionRepo *Repository) GetTransactionById(id int) (*et.Transaction, error) {
	var transaction et.Transaction
	if err := transactionRepo.DB.Preload("Transaction_detail", "deleted_at IS NULL").First(&transaction, id).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (transactionRepo *Repository) CreateTransaction(transaction *et.Transaction) error {
	if err := transactionRepo.DB.Create(&transaction).Error; err != nil {
		return err
	}

	return nil
}

func (transactionRepo *Repository) UpdateTransaction(id int, transaction *et.Transaction) error {
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
	if err := transactionRepo.DB.Where("transaction_id = ?", id).Delete(&et.TransactionDetail{}).Error; err != nil {
		return err
	}

	if err := transactionRepo.DB.Delete(&et.Transaction{}, id).Error; err != nil {
		return err
	}

	return nil
}
