package main

type DishDetails struct {
	DishName string
	Price    int
}

func NewDishDetails(dishName string, price int) *DishDetails {
	return &DishDetails{
		DishName: dishName,
		Price:    price,
	}
}

type Bill struct {
	DishDetails []*DishDetails
	PaymentMode PaymentMode
	TotalAmount int
}

func NewBill(order *Order, paymentMode PaymentMode, totalAmount int) *Bill {

	bill := &Bill{
		PaymentMode: paymentMode,
		TotalAmount: totalAmount,
		DishDetails: make([]*DishDetails, 0),
	}

	for _, dish := range order.Dish {
		bill.DishDetails = append(bill.DishDetails, NewDishDetails(dish.DishName, dish.Price))
	}

	return bill
}
