package main

import (
	"log/slog"
	"os"

	"DesignRestaurantManagementSystem/factory"
	"DesignRestaurantManagementSystem/http"
	paymentModel "DesignRestaurantManagementSystem/models/payment"
	inventoryService "DesignRestaurantManagementSystem/services/inventory"
	orderService "DesignRestaurantManagementSystem/services/order"
	paymentService "DesignRestaurantManagementSystem/services/payment"
	restaurantService "DesignRestaurantManagementSystem/services/restaurant"
)

func main() {
	// Initialize logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Initialize services
	restaurantService := restaurantService.NewRestaurantService("Mondal's Restaurant")
	inventoryService := inventoryService.NewInventoryManagementService()
	paymentFactory := factory.NewPaymentFactory()
	paymentFactory.AddPaymentService(paymentModel.CASH, paymentService.NewCashPayment())
	orderService := orderService.NewOrderManagementService(paymentFactory, inventoryService, restaurantService)

	// Initialize router with all handlers
	router := http.NewRouter(
		restaurantService,
		orderService,
		paymentFactory,
		inventoryService,
		logger,
	)

	// Setup routes
	ginRouter := router.SetupRoutes()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	logger.Info("Starting server", "port", port)
	if err := ginRouter.Run(":" + port); err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
