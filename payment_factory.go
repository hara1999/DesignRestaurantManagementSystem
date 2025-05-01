package main

import "errors"

type PaymentFactory struct {
	Method map[PaymentMode]PaymentManagementService
}

func NewPaymentFactory() *PaymentFactory {
	return &PaymentFactory{
		make(map[PaymentMode]PaymentManagementService),
	}
}

func (factory *PaymentFactory) AddPaymentService(paymentMode PaymentMode, pms PaymentManagementService) {
	factory.Method[paymentMode] = pms
}

func (factory *PaymentFactory) GetPaymentService(paymentMode PaymentMode) (PaymentManagementService, error) {

	pms, ok := factory.Method[paymentMode]
	if !ok {
		return nil, errors.New("no payment service found with provided mode")
	}
	return pms, nil
}
