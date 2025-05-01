package handlers

import (
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/dish"
	"DesignRestaurantManagementSystem/models/item"
	"DesignRestaurantManagementSystem/models/menu"

	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	interfaces.BaseHandlerInterface
	restaurantService interfaces.RestaurantServiceInterface
}

func NewMenuHandler(base interfaces.BaseHandlerInterface, service interfaces.RestaurantServiceInterface) interfaces.MenuHandlerInterface {
	return &MenuHandler{
		BaseHandlerInterface: base,
		restaurantService:    service,
	}
}

type AddDishRequest struct {
	Name     string        `json:"name" binding:"required,min=3,max=50"`
	Price    int           `json:"price" binding:"required,min=0"`
	MenuType menu.MenuType `json:"menu_type" binding:"required"`
	Items    []struct {
		Name     string `json:"name" binding:"required,min=1"`
		Quantity int    `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required,min=1"`
}

func (h *MenuHandler) AddDish(c *gin.Context) {
	var req AddDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	menu, err := h.restaurantService.GetMenuByType(req.MenuType)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	dish := dish.NewDish(req.Name, req.Price)
	for _, itemReq := range req.Items {
		item := item.NewItem(itemReq.Name, itemReq.Quantity)
		dish.AddItem(item)
	}

	menu.AddDish(dish)
	h.SuccessResponse(c, 201, dish)
}

func (h *MenuHandler) GetMenu(c *gin.Context) {
	menuType := menu.MenuType(c.Query("type"))
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

func (h *MenuHandler) UpdateDish(c *gin.Context) {
	var req AddDishRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ValidationError(c, err)
		return
	}

	menu, err := h.restaurantService.GetMenuByType(req.MenuType)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	dish, err := menu.GetDishByName(req.Name)
	if err != nil {
		h.ErrorResponse(c, 404, err.Error())
		return
	}

	dish.Price = req.Price
	// Update items if provided
	if len(req.Items) > 0 {
		dish.Item = make([]*item.Item, 0)
		for _, itemReq := range req.Items {
			item := item.NewItem(itemReq.Name, itemReq.Quantity)
			dish.AddItem(item)
		}
	}

	h.SuccessResponse(c, 200, dish)
}

func (h *MenuHandler) RemoveDish(c *gin.Context) {
	menuType := menu.MenuType(c.Query("type"))
	dishName := c.Query("name")
	if menuType == "" || dishName == "" {
		h.ErrorResponse(c, 400, "Menu type and dish name are required")
		return
	}

	menu, err := h.restaurantService.GetMenuByType(menuType)
	if err != nil {
		h.ErrorResponse(c, 400, err.Error())
		return
	}

	// Since RemoveDish is not implemented in Menu, we'll remove it manually
	found := false
	for i, d := range menu.Dish {
		if d.DishName == dishName {
			menu.Dish = append(menu.Dish[:i], menu.Dish[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		h.ErrorResponse(c, 404, "Dish not found")
		return
	}

	h.SuccessResponse(c, 200, gin.H{"message": "Dish removed successfully"})
}
