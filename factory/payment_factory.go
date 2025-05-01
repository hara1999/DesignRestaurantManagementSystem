package factory

import (
	"errors"

	paymentModel "DesignRestaurantManagementSystem/models/payment"
	paymentService "DesignRestaurantManagementSystem/services/payment"
)

type PaymentFactory struct {
	Method map[paymentModel.PaymentMode]paymentService.PaymentManagementService
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{
		make(map[paymentModel.PaymentMode]paymentService.PaymentManagementService),
	}
}

func (factory *PaymentFactory) AddPaymentService(paymentMode paymentModel.PaymentMode, pms paymentService.PaymentManagementService) {
	factory.Method[paymentMode] = pms
}

func (factory *PaymentFactory) GetPaymentService(paymentMode paymentModel.PaymentMode) (paymentService.PaymentManagementService, error) {

	pms, ok := factory.Method[paymentMode]
	if !ok {
		return nil, errors.New("no payment service found with provided mode")
	}
	return pms, nil
}
