package main

type Item struct {
	ItemName string
	Quantity int
}

func NewItem(itemName string, quantity int) *Item {
	return &Item{
		ItemName: itemName,
		Quantity: quantity,
	}
}
