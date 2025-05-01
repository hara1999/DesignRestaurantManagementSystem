package interfaces

import "github.com/gin-gonic/gin"

type BaseHandlerInterface interface {
	SuccessResponse(c *gin.Context, status int, data interface{})
	ErrorResponse(c *gin.Context, status int, message string)
	ValidationError(c *gin.Context, err error)
	InternalServerError(c *gin.Context, err error)
}

type RestaurantHandlerInterface interface {
	AddTable(c *gin.Context)
	GetTables(c *gin.Context)
	ReserveTable(c *gin.Context)
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
}

type MenuHandlerInterface interface {
	AddDish(c *gin.Context)
	GetMenu(c *gin.Context)
	UpdateDish(c *gin.Context)
	RemoveDish(c *gin.Context)
}
