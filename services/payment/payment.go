package payment

import (
	billModel "DesignRestaurantManagementSystem/models/bill"
	orderModel "DesignRestaurantManagementSystem/models/order"
	paymentModel "DesignRestaurantManagementSystem/models/payment"
	"fmt"

	"DesignRestaurantManagementSystem/interfaces"
)

// concrete implementation of the interface
type CashPayment struct {
}

func NewCashPayment() interfaces.PaymentManagementServiceInterface {
	return &CashPayment{}
}

func (payment *CashPayment) ProcessPayment(order *orderModel.Order) (*billModel.Bill, error) {

	totalAmount := 0
	for _, dish := range order.Dish {
		dishCost, err := order.GetDishPriceByName(dish.DishName)
		if err != nil {
			return nil, fmt.Errorf("payment failed due to: %w", err)
		}
		totalAmount = totalAmount + dishCost
	}

	// process payment according to the respective logic for this mode

	bill := billModel.NewBill(order, paymentModel.CASH, totalAmount)
	return bill, nil
}
