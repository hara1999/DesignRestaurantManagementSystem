package main

import (
	"errors"
)

type OrderStatus string

const (
	OnProcess OrderStatus = "OnProcess"
	Processed OrderStatus = "Processed"
)

type Order struct {
	OrderID string
	Dish    []*Dish
	Status  OrderStatus
}

func NewOrder(orderID string, dish []*Dish) *Order {
	return &Order{
		OrderID: orderID,
		Dish:    dish,
		Status:  OnProcess,
	}
}

func (o *Order) AddDish(dish *Dish) {
	o.Dish = append(o.Dish, dish)
}

func (o *Order) GetDishByName(dishName string) (*Dish, error) {
	for _, dish := range o.Dish {
		if dish.DishName == dishName {
			return dish, nil
		}
	}
	return nil, errors.New("no dish found")
}

func (o *Order) UpdateStatus(status OrderStatus) {
	o.Status = status
}

func (o *Order) GetDishPriceByName(dishName string) (int, error) {
	dish, err := o.GetDishByName(dishName)
	if err != nil {
		return 0, err
	}
	return dish.Price, nil
}
