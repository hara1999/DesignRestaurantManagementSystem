package dish

import (
	itemModel "DesignRestaurantManagementSystem/models/item"
)

type AddDishRequest struct {
	Name     string                     `json:"name"`
	Price    int                        `json:"price"`
	MenuType string                     `json:"menu_type"`
	Items    []itemModel.AddItemRequest `json:"items"`
}
