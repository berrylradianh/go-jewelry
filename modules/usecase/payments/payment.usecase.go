package payments

import (
	e "github.com/berrylradianh/go-jewelry/modules/entity"
	rp "github.com/berrylradianh/go-jewelry/modules/repository/payments"
)

type Usecase struct {
	Repository rp.Repository
}

func (paymentUsecase *Usecase) GetAllPayments() (*[]e.Payment, error) {
	payments, err := paymentUsecase.Repository.GetAllPayments()
	return payments, err
}

func (paymentUsecase *Usecase) GetPaymentById(id int) (*e.Payment, error) {
	payment, err := paymentUsecase.Repository.GetPaymentById(id)
	return payment, err
}

func (paymentUsecase *Usecase) CreatePayment(payment *e.Payment) error {
	err := paymentUsecase.Repository.CreatePayment(payment)
	if err != nil {
		return err
	}

	return nil
}

func (paymentUsecase *Usecase) UpdatePayment(id int, payment *e.Payment) error {
	result := paymentUsecase.Repository.UpdatePayment(id, payment)
	return result
}

func (paymentUsecase *Usecase) DeletePayment(id int) error {
	err := paymentUsecase.Repository.DeletePayment(id)
	return err
}
