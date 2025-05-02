package interfaces

import "github.com/gin-gonic/gin"

type BaseHandlerInterface interface {
	SuccessResponse(c *gin.Context, status int, data interface{})
	ErrorResponse(c *gin.Context, status int, message string)
	ValidationError(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

type RestaurantHandlerInterface interface {
	AddMenu(c *gin.Context)
	AddTable(c *gin.Context)
	AddDish(c *gin.Context)
	GetMenuByType(c *gin.Context)
	GetTableByID(c *gin.Context)
	GetTables(c *gin.Context)
	ReserveTable(c *gin.Context)
	UpdateDish(c *gin.Context)
	RemoveDish(c *gin.Context)
}

type OrderHandlerInterface interface {
	PlaceOrder(c *gin.Context)
	GetOrderStatus(c *gin.Context)
	CancelOrder(c *gin.Context)
	GetOrderHistory(c *gin.Context)
}

type PaymentHandlerInterface interface {
	ProcessPayment(c *gin.Context)
	GetPaymentStatus(c *gin.Context)
	RefundPayment(c *gin.Context)
}

type InventoryHandlerInterface interface {
	AddItem(c *gin.Context)
	GetInventory(c *gin.Context)
	UpdateItemQuantity(c *gin.Context)
	RemoveItem(c *gin.Context)
	GetItemByStatus(c *gin.Context)
	RemoveExpiredItems(c *gin.Context)
}
