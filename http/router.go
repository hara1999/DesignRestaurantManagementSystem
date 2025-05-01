package http

import (
	"DesignRestaurantManagementSystem/factory"
	"DesignRestaurantManagementSystem/handlers"
	"DesignRestaurantManagementSystem/interfaces"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Router struct {
	restaurantHandler interfaces.RestaurantHandlerInterface
	orderHandler      interfaces.OrderHandlerInterface
	paymentHandler    interfaces.PaymentHandlerInterface
	inventoryHandler  interfaces.InventoryHandlerInterface
	menuHandler       interfaces.MenuHandlerInterface
}

func NewRouter(
	restaurantService interfaces.RestaurantServiceInterface,
	orderService interfaces.OrderManagementServiceInterface,
	paymentFactory *factory.PaymentFactory,
	inventoryService interfaces.InventoryManagementServiceInterface,
	logger *slog.Logger,
) *Router {
	baseHandler := handlers.NewBaseHandler(logger)

	return &Router{
		restaurantHandler: handlers.NewRestaurantHandler(baseHandler, restaurantService),
		orderHandler:      handlers.NewOrderHandler(baseHandler, orderService),
		paymentHandler:    handlers.NewPaymentHandler(baseHandler, paymentFactory),
		inventoryHandler:  handlers.NewInventoryHandler(baseHandler, inventoryService),
		menuHandler:       handlers.NewMenuHandler(baseHandler, restaurantService),
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
