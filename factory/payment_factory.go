package factory

import (
	"errors"

	"DesignRestaurantManagementSystem/interfaces"
	paymentModel "DesignRestaurantManagementSystem/models/payment"
)

type PaymentFactory struct {
	Method map[paymentModel.PaymentMode]interfaces.PaymentManagementServiceInterface
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{
		make(map[paymentModel.PaymentMode]interfaces.PaymentManagementServiceInterface),
	}
}

func (factory *PaymentFactory) AddPaymentService(paymentMode paymentModel.PaymentMode, pms interfaces.PaymentManagementServiceInterface) {
	factory.Method[paymentMode] = pms
}

func (factory *PaymentFactory) GetPaymentService(paymentMode paymentModel.PaymentMode) (interfaces.PaymentManagementServiceInterface, error) {

	pms, ok := factory.Method[paymentMode]
	if !ok {
		return nil, errors.New("no payment service found with provided mode")
	}
	return pms, nil
}
