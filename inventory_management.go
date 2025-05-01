package main

import "errors"

type InventoryManagementService struct {
	Item []*Item
}

func NewInventoryManagementService() *InventoryManagementService {
	return &InventoryManagementService{
		Item: make([]*Item, 0),
	}
}

func (ims *InventoryManagementService) AddItem(item *Item) {
	ims.Item = append(ims.Item, item)
}

func (ims *InventoryManagementService) GetItemByName(itemName string) (*Item, error) {
	for _, item := range ims.Item {
		if item.ItemName == itemName {
			return item, nil
		}
	}
	return nil, errors.New("no item found")
}

func (ims *InventoryManagementService) GetItemsByName(itemNames []string) ([]*Item, []string) {

	var foundItem []*Item
	var notFoundItem []string

	for _, itemName := range itemNames {
		item, err := ims.GetItemByName(itemName)
		if err != nil {
			notFoundItem = append(notFoundItem, item.ItemName)
		} else {
			foundItem = append(foundItem, item)
		}
	}
	return foundItem, notFoundItem
}
