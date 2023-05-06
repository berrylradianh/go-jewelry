package payments

import (
	ep "github.com/berrylradianh/go-jewelry/modules/entity/payments"
	rp "github.com/berrylradianh/go-jewelry/modules/repository/payments"
)

type Usecase struct {
	Repository rp.Repository
}

func (paymentUsecase *Usecase) GetAllPayments() (*[]ep.Payment, error) {
	payments, err := paymentUsecase.Repository.GetAllPayments()
	return payments, err
}

func (paymentUsecase *Usecase) GetPaymentById(id int) (*ep.Payment, error) {
	payment, err := paymentUsecase.Repository.GetPaymentById(id)
	return payment, err
}

func (paymentUsecase *Usecase) CreatePayment(payment *ep.Payment) error {
	err := paymentUsecase.Repository.CreatePayment(payment)
	if err != nil {
		return err
	}

	return nil
}

func (paymentUsecase *Usecase) UpdatePayment(id int, payment *ep.Payment) error {
	result := paymentUsecase.Repository.UpdatePayment(id, payment)
	return result
}

func (paymentUsecase *Usecase) DeletePayment(id int) error {
	err := paymentUsecase.Repository.DeletePayment(id)
	return err
}
