package dish

import (
	"DesignRestaurantManagementSystem/models/item"
	"errors"
)

type DishStatus string

const (
	DISH_AVAILABLE     DishStatus = "DISH_AVAILABLE"
	DISH_NOT_AVAILABLE DishStatus = "DISH_NOT_AVAILABLE"
)

type Dish struct {
	DishName     string
	Price        int
	Item         []*item.Item
	Availability DishStatus
}

func NewDish(dishName string, price int) *Dish {
	return &Dish{
		DishName:     dishName,
		Price:        price,
		Item:         make([]*item.Item, 0),
		Availability: DISH_AVAILABLE,
	}
}

func (d *Dish) AddItem(item *item.Item) {
	d.Item = append(d.Item, item)
}

func (d *Dish) GetItemByName(itemName string) (*item.Item, error) {
	for _, item := range d.Item {
		if item.ItemName == itemName {
			return item, nil
		}
	}
	return nil, errors.New("no item found")
}

func (d *Dish) GetItemsWithNames() ([]*item.Item, []string, error) {
	var itemNames []string
	var items []*item.Item
	for _, item := range d.Item {
		itemNames = append(itemNames, item.ItemName)
		items = append(items, item)
	}

	if len(itemNames) > 0 {
		return items, itemNames, nil
	}
	return nil, nil, errors.New("no item found")
}
