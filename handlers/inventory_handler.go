package handlers

import (
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/item"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	interfaces.BaseHandlerInterface
	inventoryService interfaces.InventoryManagementServiceInterface
}

func NewInventoryHandler(base interfaces.BaseHandlerInterface, service interfaces.InventoryManagementServiceInterface) interfaces.InventoryHandlerInterface {
	return &InventoryHandler{
		BaseHandlerInterface: base,
		inventoryService:     service,
	}
}

func (h *InventoryHandler) AddItem(c *gin.Context) {
	var req item.AddItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	item, err := h.inventoryService.AddItem(req.Name, req.Quantity)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}
	h.SuccessResponse(c, 201, item)
}

func (h *InventoryHandler) GetInventory(c *gin.Context) {
	items := h.inventoryService.GetItems()
	h.SuccessResponse(c, 200, items)
}

func (h *InventoryHandler) UpdateItemQuantity(c *gin.Context) {
	var req item.AddItemRequest
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

func (h *InventoryHandler) GetItemByStatus(c *gin.Context) {
	status := c.Query("status")
	if status == "" {
		h.ErrorResponse(c, 400, "Status is required")
		return
	}

	items, err := h.inventoryService.GetItemByStatus(status)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}
	h.SuccessResponse(c, 200, items)
}

func (h *InventoryHandler) RemoveExpiredItems(c *gin.Context) {
	expiredItems := h.inventoryService.RemoveExpiredItems()
	h.SuccessResponse(c, 200, expiredItems)
}
