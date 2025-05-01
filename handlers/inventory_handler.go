package handlers

import (
	"DesignRestaurantManagementSystem/models/item"
	"DesignRestaurantManagementSystem/services/inventory"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	*BaseHandler
	inventoryService *inventory.InventoryManagementService
}

func NewInventoryHandler(base *BaseHandler, service *inventory.InventoryManagementService) *InventoryHandler {
	return &InventoryHandler{
		BaseHandler:      base,
		inventoryService: service,
	}
}

type AddItemRequest struct {
	Name     string `json:"name" binding:"required,min=1"`
	Quantity int    `json:"quantity" binding:"required,min=1"`
}

func (h *InventoryHandler) AddItem(c *gin.Context) {
	var req AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	item := item.NewItem(req.Name, req.Quantity)
	h.inventoryService.AddItem(item)
	h.SuccessResponse(c, 201, item)
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	items := h.inventoryService.GetItems()
	h.SuccessResponse(c, 200, items)
}

func (h *InventoryHandler) UpdateItemQuantity(c *gin.Context) {
	var req AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.inventoryService.UpdateItemQuantity(req.Name, req.Quantity)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Item quantity updated successfully"})
}

func (h *InventoryHandler) RemoveItem(c *gin.Context) {
	itemName := c.Query("name")
	if itemName == "" {
		h.ErrorResponse(c, 400, "Item name is required")
		return
	}

	err := h.inventoryService.RemoveItem(itemName)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Item removed successfully"})
}
