package handlers

import (
	"DesignRestaurantManagementSystem/factory"
	"DesignRestaurantManagementSystem/services/inventory"
	"DesignRestaurantManagementSystem/services/order"
	"DesignRestaurantManagementSystem/services/restaurant"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Router struct {
	restaurantHandler *RestaurantHandler
	orderHandler      *OrderHandler
	paymentHandler    *PaymentHandler
	inventoryHandler  *InventoryHandler
	menuHandler       *MenuHandler
}

func NewRouter(
	restaurantService *restaurant.RestaurantService,
	orderService *order.OrderManagementService,
	paymentFactory *factory.PaymentFactory,
	inventoryService *inventory.InventoryManagementService,
	logger *slog.Logger,
) *Router {
	baseHandler := NewBaseHandler(logger)

	return &Router{
		restaurantHandler: NewRestaurantHandler(baseHandler, restaurantService),
		orderHandler:      NewOrderHandler(baseHandler, orderService),
		paymentHandler:    NewPaymentHandler(baseHandler, paymentFactory),
		inventoryHandler:  NewInventoryHandler(baseHandler, inventoryService),
		menuHandler:       NewMenuHandler(baseHandler, restaurantService),
	}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Restaurant routes
	restaurantGroup := router.Group("/restaurant")
	{
		restaurantGroup.GET("/tables", r.restaurantHandler.GetTables)
		restaurantGroup.POST("/tables/add", r.restaurantHandler.AddTable)
		restaurantGroup.POST("/tables/reserve", r.restaurantHandler.ReserveTable)
	}

	// Order routes
	orderGroup := router.Group("/orders")
	{
		orderGroup.POST("", r.orderHandler.PlaceOrder)
		orderGroup.GET("/status", r.orderHandler.GetOrderStatus)
		orderGroup.POST("/cancel", r.orderHandler.CancelOrder)
	}

	// Payment routes
	paymentGroup := router.Group("/payments")
	{
		paymentGroup.POST("", r.paymentHandler.ProcessPayment)
		paymentGroup.GET("/status", r.paymentHandler.GetPaymentStatus)
		paymentGroup.POST("/refund", r.paymentHandler.RefundPayment)
	}

	// Inventory routes
	inventoryGroup := router.Group("/inventory")
	{
		inventoryGroup.GET("", r.inventoryHandler.GetInventory)
		inventoryGroup.POST("/add", r.inventoryHandler.AddItem)
		inventoryGroup.POST("/update", r.inventoryHandler.UpdateItemQuantity)
		inventoryGroup.POST("/remove", r.inventoryHandler.RemoveItem)
	}

	// Menu routes
	menuGroup := router.Group("/menu")
	{
		menuGroup.GET("", r.menuHandler.GetMenu)
		menuGroup.POST("/add", r.menuHandler.AddDish)
		menuGroup.POST("/update", r.menuHandler.UpdateDish)
		menuGroup.POST("/remove", r.menuHandler.RemoveDish)
	}

	return router
}
