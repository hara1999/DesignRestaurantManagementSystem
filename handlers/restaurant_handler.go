package handlers

import (
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/table"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
	interfaces.BaseHandlerInterface
	restaurantService interfaces.RestaurantServiceInterface
}

func NewRestaurantHandler(base interfaces.BaseHandlerInterface, service interfaces.RestaurantServiceInterface) *RestaurantHandler {
	return &RestaurantHandler{
		BaseHandlerInterface: base,
		restaurantService:    service,
	}
}

type CreateRestaurantRequest struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
}

type AddTableRequest struct {
	TableID string `json:"table_id" binding:"required,min=1,max=10"`
}

func (h *RestaurantHandler) AddTable(c *gin.Context) {
	var req AddTableRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	table := table.NewTable(req.TableID)
	h.restaurantService.AddTable(table)

	h.SuccessResponse(c, 201, table)
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

// func (h *RestaurantHandler) GetRestaurantDetails(c *gin.Context) {
// 	details, err := h.restaurantService.GetDetails()
// 	if err != nil {
// 		h.ErrorResponse(c, 400, err.Error())
// 		return
// 	}

// 	h.SuccessResponse(c, 200, details)
// }
