package interfaces

import (
	billModel "DesignRestaurantManagementSystem/models/bill"
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
	AddItem(item *itemModel.Item)
	GetItemByName(itemName string) (*itemModel.Item, error)
	GetItemsByName(itemNames []string) ([]*itemModel.Item, []string)
	GetItems() []*itemModel.Item
	UpdateItemQuantity(itemName string, quantity int) error
	RemoveItem(itemName string) error
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
	AddMenu(menu *menuModel.Menu)
	GetMenuByType(menuType menuModel.MenuType) (*menuModel.Menu, error)
	AddTable(table *tableModel.Table)
	GetTableByID(tableID string) (*tableModel.Table, error)
	GetTables() ([]*tableModel.Table, error)
	ReserveTable(tableIDs []string) error
}
