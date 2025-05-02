package interfaces

import (
	billModel "DesignRestaurantManagementSystem/models/bill"
	dishModel "DesignRestaurantManagementSystem/models/dish"
	employeeModel "DesignRestaurantManagementSystem/models/employee"
	itemModel "DesignRestaurantManagementSystem/models/item"
	menuModel "DesignRestaurantManagementSystem/models/menu"
	orderModel "DesignRestaurantManagementSystem/models/order"
	paymentModel "DesignRestaurantManagementSystem/models/payment"
	tableModel "DesignRestaurantManagementSystem/models/table"
)

type EmployeeManagementServiceInterface interface {
	EnrollEmployee(emp *employeeModel.Employee)
	GetEmployeeByName(name string) (*employeeModel.Employee, error)
	GetEmployeesByDesignation(designation employeeModel.Designation) ([]*employeeModel.Employee, error)
	AssignShift(employeeName string, shift employeeModel.WorkShift) error
}

type InventoryManagementServiceInterface interface {
	AddItem(itemName string, quantity int) (*itemModel.Item, error)
	GetItemByName(itemName string) (*itemModel.Item, error)
	GetItemsByName(itemNames []string) ([]*itemModel.Item, []string)
	GetItems() []*itemModel.Item
	UpdateItemQuantity(itemName string, quantity int) error
	RemoveItem(itemName string) error
	GetItemByStatus(status string) ([]*itemModel.Item, error)
	RemoveExpiredItems() []*itemModel.Item
}

type OrderManagementServiceInterface interface {
	PlaceOrder(dishName []string, menuType menuModel.MenuType, paymentMode paymentModel.PaymentMode) (*billModel.Bill, error)
	GetOrderStatus(orderID string) (orderModel.OrderStatus, error)
	CancelOrder(orderID string) error
	GetOrderHistory() ([]*orderModel.Order, error)
}

type PaymentManagementServiceInterface interface {
	ProcessPayment(order *orderModel.Order) (*billModel.Bill, error)
}

type RestaurantServiceInterface interface {
	AddMenu(menuReq menuModel.AddMenuRequest) error
	AddTable(tableReq tableModel.AddTableRequest) error
	AddDish(dishReq dishModel.AddDishRequest) error
	GetMenuByType(menuType string) (*menuModel.Menu, error)
	GetTableByID(tableID string) (*tableModel.Table, error)
	GetTables() ([]*tableModel.Table, error)
	ReserveTable(tableIDs []string) error
	UpdateDish(dishReq dishModel.AddDishRequest) error
	RemoveDish(menuType string, dishName string) error
}
