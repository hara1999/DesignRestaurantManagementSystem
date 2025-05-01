package main

import (
	"log/slog"
)

func createMenu(restaurant *Restaurant) *Menu {
	menu := NewMenu(MenuType(Starter))

	dish1 := NewDish("chicken tikka", 50)
	dish1.AddItem(NewItem("onion", 7))

	menu.AddDish(dish1)
	// menu.AddDish(NewDish("mutton kebab", 100))

	restaurant.AddMenu(menu)

	return menu
}

func main() {
	restaurant := NewRestaurant("Mondal's R")

	createMenu(restaurant)

	restaurant.AddTable(NewTable("1"))
	restaurant.AddTable(NewTable("2"))

	inventoryManagementService := NewInventoryManagementService()
	inventoryManagementService.AddItem(NewItem("onion", 4))

	paymentFactory := NewPaymentFactory()
	paymentFactory.AddPaymentService(CASH, NewCashPayment())

	orderService := NewOrderManagementService(paymentFactory, inventoryManagementService, restaurant)

	bill, err := orderService.OrderFood([]string{"chicken tikka"}, Starter, CASH)
	if err != nil {
		slog.Error("can not make the order", "error", err)
	} else {
		slog.Info("Successfully ordered food: ", "paymentmode", bill.PaymentMode, "totalAmount", bill.TotalAmount)
	}

	// err = restaurant.ReserveTable([]string{"1", "2"})
	// if err != nil {
	// 	slog.Error("can not book the table", "error", err)
	// }

	// err = restaurant.ReserveTable([]string{"2"})
	// if err != nil {
	// 	slog.Error("can not book the table", "error", err)
	// }
}
