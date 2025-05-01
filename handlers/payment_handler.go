package handlers

import (
	"DesignRestaurantManagementSystem/factory"
	"DesignRestaurantManagementSystem/models/order"
	"DesignRestaurantManagementSystem/models/payment"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	*BaseHandler
	paymentFactory *factory.PaymentFactory
}

func NewPaymentHandler(base *BaseHandler, paymentFactory *factory.PaymentFactory) *PaymentHandler {
	return &PaymentHandler{
		BaseHandler:    base,
		paymentFactory: paymentFactory,
	}
}

type ProcessPaymentRequest struct {
	OrderID string              `json:"order_id" binding:"required"`
	Amount  float64             `json:"amount" binding:"required,min=0"`
	Mode    payment.PaymentMode `json:"payment_mode" binding:"required"`
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req ProcessPaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	// Create a simple order for processing
	mockOrder := order.NewOrder(req.OrderID, nil)

	// Process the payment through the service
	paymentService, err := h.paymentFactory.GetPaymentService(req.Mode)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	bill, err := paymentService.Process(mockOrder)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, bill)
}

// Note: The PaymentManagementService interface only includes Process method
// For completeness, we'd need to extend the interface or create a repository
// The following methods would need actual implementation on the backend

func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	paymentID := c.Query("payment_id")
	if paymentID == "" {
		h.ErrorResponse(c, 400, "Payment ID is required")
		return
	}

	// This would call the actual payment service
	h.SuccessResponse(c, 200, gin.H{
		"payment_id": paymentID,
		"status":     "COMPLETED", // Mocked status
	})
}

func (h *PaymentHandler) RefundPayment(c *gin.Context) {
	paymentID := c.Query("payment_id")
	if paymentID == "" {
		h.ErrorResponse(c, 400, "Payment ID is required")
		return
	}

	// This would call the actual refund implementation
	h.SuccessResponse(c, 200, gin.H{"message": "Payment refunded successfully"})
}
