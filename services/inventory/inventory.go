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

func (ims *InventoryManagementService) AddItem(itemName string, quantity int) (*itemModel.Item, error) {
	var item *itemModel.Item
	_, err := ims.GetItemByName(itemName)
	if err != nil {
		item = itemModel.NewItem(itemName, quantity)
		ims.Items = append(ims.Items, item)
		return item, nil
	}
	return nil, errors.New("item already exists")
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

	if item.Status == itemModel.ItemStatusExpired {
		ims.Items = append(ims.Items, itemModel.NewItem(itemName, 0))
	}

	item.Quantity = item.Quantity + quantity
	return nil
}

func (ims *InventoryManagementService) RemoveItem(itemName string) error {
	item, err := ims.GetItemByName(itemName)
	if err != nil {
		return fmt.Errorf("RemoveItem: %w", err)
	}
	item.Status = itemModel.ItemStatusExpired
	return nil
}

func (ims *InventoryManagementService) GetItemByStatus(status string) ([]*itemModel.Item, error) {
	var items []*itemModel.Item
	for _, item := range ims.Items {
		if item.Status == itemModel.ItemStatus(status) {
			items = append(items, item)
		}
	}
	return items, nil
}

func (ims *InventoryManagementService) RemoveExpiredItems() []*itemModel.Item {

	expiredItems := make([]*itemModel.Item, 0)

	for i, item := range ims.Items {
		if item.Status == itemModel.ItemStatusExpired {
			expiredItems = append(expiredItems, item)
			ims.Items = append(ims.Items[:i], ims.Items[i+1:]...)
		}
	}

	return expiredItems
}
