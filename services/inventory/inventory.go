package inventory

import (
	"DesignRestaurantManagementSystem/interfaces"
	itemModel "DesignRestaurantManagementSystem/models/item"
	"errors"
	"fmt"
)

type InventoryManagementService struct {
	Items []*itemModel.Item
}

func NewInventoryManagementService() interfaces.InventoryManagementServiceInterface {
	return &InventoryManagementService{
		Items: make([]*itemModel.Item, 0),
	}
}

func (ims *InventoryManagementService) AddItem(item *itemModel.Item) {
	ims.Items = append(ims.Items, item)
}

func (ims *InventoryManagementService) GetItemByName(itemName string) (*itemModel.Item, error) {
	for _, item := range ims.Items {
		if item.ItemName == itemName {
			return item, nil
		}
	}
	return nil, errors.New("no item found")
}

func (ims *InventoryManagementService) GetItemsByName(itemNames []string) ([]*itemModel.Item, []string) {

	var foundItems []*itemModel.Item
	var notFoundItems []string

	for _, itemName := range itemNames {
		item, err := ims.GetItemByName(itemName)
		if err != nil {
			notFoundItems = append(notFoundItems, itemName)
		} else {
			foundItems = append(foundItems, item)
		}
	}
	return foundItems, notFoundItems
}

func (ims *InventoryManagementService) GetItems() []*itemModel.Item {
	return ims.Items
}

func (ims *InventoryManagementService) UpdateItemQuantity(itemName string, quantity int) error {
	item, err := ims.GetItemByName(itemName)
	if err != nil {
		return fmt.Errorf("UpdateItemQuantity: %w", err)
	}
	item.Quantity = quantity
	return nil
}

func (ims *InventoryManagementService) RemoveItem(itemName string) error {
	for i, item := range ims.Items {
		if item.ItemName == itemName {
			ims.Items = append(ims.Items[:i], ims.Items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("RemoveItem: %w", errors.New("item not found"))
}
