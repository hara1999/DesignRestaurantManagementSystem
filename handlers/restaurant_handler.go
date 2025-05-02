package handlers

import (
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/dish"
	"DesignRestaurantManagementSystem/models/menu"
	"DesignRestaurantManagementSystem/models/table"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	interfaces.BaseHandlerInterface
	restaurantService interfaces.RestaurantServiceInterface
}

func NewRestaurantHandler(base interfaces.BaseHandlerInterface, service interfaces.RestaurantServiceInterface) interfaces.RestaurantHandlerInterface {
	return &RestaurantHandler{
		BaseHandlerInterface: base,
		restaurantService:    service,
	}
}

func (h *RestaurantHandler) AddMenu(c *gin.Context) {
	var req menu.AddMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.restaurantService.AddMenu(req)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 201, req)
}

func (h *RestaurantHandler) AddDish(c *gin.Context) {
	var req dish.AddDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.restaurantService.AddDish(req)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 201, req)
}

func (h *RestaurantHandler) AddTable(c *gin.Context) {
	var req table.AddTableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.restaurantService.AddTable(req)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 201, req)
}

func (h *RestaurantHandler) GetTables(c *gin.Context) {
	tables, err := h.restaurantService.GetTables()
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, tables)
}

func (h *RestaurantHandler) ReserveTable(c *gin.Context) {
	var tableIDs []string
	if err := c.ShouldBindJSON(&tableIDs); err != nil {
		h.ValidationError(c, err)
		return
	}

	if len(tableIDs) == 0 {
		h.ErrorResponse(c, 400, "At least one table ID is required")
		return
	}

	err := h.restaurantService.ReserveTable(tableIDs)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Tables reserved successfully"})
}

func (h *RestaurantHandler) GetMenuByType(c *gin.Context) {
	menuType := c.Query("type")
	if menuType == "" {
		h.ErrorResponse(c, 400, "Menu type is required")
		return
	}

	menu, err := h.restaurantService.GetMenuByType(menuType)
	if err != nil {
		h.ErrorResponse(c, 404, err.Error())
		return
	}

	h.SuccessResponse(c, 200, menu)
}

func (h *RestaurantHandler) UpdateDish(c *gin.Context) {
	var req dish.AddDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	err := h.restaurantService.UpdateDish(req)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, req)
}

func (h *RestaurantHandler) RemoveDish(c *gin.Context) {
	menuType := c.Query("type")
	dishName := c.Query("name")
	if menuType == "" || dishName == "" {
		h.ErrorResponse(c, 400, "Menu type and dish name are required")
		return
	}

	err := h.restaurantService.RemoveDish(menuType, dishName)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Dish removed successfully"})
}

func (h *RestaurantHandler) GetTableByID(c *gin.Context) {
	tableID := c.Query("id")
	if tableID == "" {
		h.ErrorResponse(c, 400, "Table ID is required")
		return
	}

	table, err := h.restaurantService.GetTableByID(tableID)
	if err != nil {
		h.ErrorResponse(c, 404, err.Error())
		return
	}

	h.SuccessResponse(c, 200, table)
}
