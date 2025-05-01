package main

import "fmt"

type PaymentMode string

const (
	CASH   PaymentMode = "CASH"
	MOBILE PaymentMode = "MOBILE"
)

type PaymentManagementService interface {
	Payment(order *Order) (*Bill, error)
}

type CashPayment struct {
}

func NewCashPayment() PaymentManagementService {
	return &CashPayment{}
}

func (payment *CashPayment) Payment(order *Order) (*Bill, error) {

	totalAmount := 0
	for _, dish := range order.Dish {
		dishCost, err := order.GetDishPriceByName(dish.DishName)
		if err != nil {
			return nil, fmt.Errorf("payment failed due to: %w", err)
		}
		totalAmount = totalAmount + dishCost
	}

	bill := NewBill(order, CASH, totalAmount)
	return bill, nil
}
