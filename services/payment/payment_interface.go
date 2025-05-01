package payment

import (
	"DesignRestaurantManagementSystem/models/bill"
	"DesignRestaurantManagementSystem/models/order"
)

type PaymentManagementService interface {
	Process(order *order.Order) (*bill.Bill, error)
}
