package menu

import (
	dishModel "DesignRestaurantManagementSystem/models/dish"
)

type AddMenuRequest struct {
	MenuType string                     `json:"menu_type"`
	Dish     []dishModel.AddDishRequest `json:"dish"`
}
