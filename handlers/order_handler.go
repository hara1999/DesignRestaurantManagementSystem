package handlers

import (
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/menu"
	"DesignRestaurantManagementSystem/models/payment"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	interfaces.BaseHandlerInterface
	orderService interfaces.OrderManagementServiceInterface
}

func NewOrderHandler(base interfaces.BaseHandlerInterface, service interfaces.OrderManagementServiceInterface) interfaces.OrderHandlerInterface {
	return &OrderHandler{
		BaseHandlerInterface: base,
		orderService:         service,
	}
}

type PlaceOrderRequest struct {
	DishNames   []string            `json:"dish_names" binding:"required,min=1"`
	MenuType    menu.MenuType       `json:"menu_type" binding:"required"`
	PaymentMode payment.PaymentMode `json:"payment_mode" binding:"required"`
}

type OrderStatusRequest struct {
	OrderID string `json:"order_id" binding:"required"`
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	bill, err := h.orderService.PlaceOrder(req.DishNames, req.MenuType, req.PaymentMode)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 201, bill)
}

func (h *OrderHandler) GetOrderStatus(c *gin.Context) {
	var req OrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	status, err := h.orderService.GetOrderStatus(req.OrderID)
	if err != nil {
		h.ErrorResponse(c, 404, err.Error())
		return
	}

	h.SuccessResponse(c, 200, status)
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	var req OrderStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.orderService.CancelOrder(req.OrderID)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Order cancelled successfully"})
}

func (h *OrderHandler) GetOrderHistory(c *gin.Context) {
	history, err := h.orderService.GetOrderHistory()
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, history)
}
