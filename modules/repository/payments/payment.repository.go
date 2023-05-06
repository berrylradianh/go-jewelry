package payments

import (
	"fmt"

	ep "github.com/berrylradianh/go-jewelry/modules/entity/payments"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (paymentRepo *Repository) GetAllPayments() (*[]ep.Payment, error) {
	var payments []ep.Payment
	if err := paymentRepo.DB.Find(&payments).Error; err != nil {
		return nil, err
	}

	return &payments, nil
}

func (paymentRepo *Repository) GetPaymentById(id int) (*ep.Payment, error) {
	var payment ep.Payment
	if err := paymentRepo.DB.First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}

func (paymentRepo *Repository) CreatePayment(payment *ep.Payment) error {
	if err := paymentRepo.DB.Create(&payment).Error; err != nil {
		return err
	}

	return nil
}

func (paymentRepo *Repository) UpdatePayment(id int, payment *ep.Payment) error {
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
	if err := paymentRepo.DB.Delete(&ep.Payment{}, id).Error; err != nil {
		return err
	}

	return nil
}
