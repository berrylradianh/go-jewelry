package payments

import (
	"fmt"

	e "github.com/berrylradianh/go-jewelry/modules/entity"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (paymentRepo *Repository) GetAllPayments() (*[]e.Payment, error) {
	var payments []e.Payment
	if err := paymentRepo.DB.Preload("Transaction", "deleted_at IS NULL").Find(&payments).Error; err != nil {
		return nil, err
	}

	return &payments, nil
}

func (paymentRepo *Repository) GetPaymentById(id int) (*e.Payment, error) {
	var payment e.Payment
	if err := paymentRepo.DB.Preload("Transaction", "deleted_at IS NULL").First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (paymentRepo *Repository) CreatePayment(payment *e.Payment) error {
	if err := paymentRepo.DB.Create(&payment).Error; err != nil {
		return err
	}

	return nil
}

func (paymentRepo *Repository) UpdatePayment(id int, payment *e.Payment) error {
	result := paymentRepo.DB.Model(&payment).Where("id = ?", id).Omit("UpdatedAt").Updates(&payment)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("nothing updated")
	}

	return nil
}

func (paymentRepo *Repository) DeletePayment(id int) error {
	if err := paymentRepo.DB.Delete(&e.Payment{}, id).Error; err != nil {
		return err
	}

	return nil
}
