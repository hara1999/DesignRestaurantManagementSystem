package main

import (
	"errors"
	"fmt"
)

type OrderManagementService struct {
	paymentFactory *PaymentFactory
	ims            *InventoryManagementService
	restaurant     *Restaurant
	orders         []*Order
}

func NewOrderManagementService(paymentFactory *PaymentFactory, ims *InventoryManagementService, restaurant *Restaurant) *OrderManagementService {
	return &OrderManagementService{
		paymentFactory: paymentFactory,
		ims:            ims,
		restaurant:     restaurant,
		orders:         make([]*Order, 0),
	}
}

func checkQuantity(available []*Item, needed []*Item) error {

	for _, aItem := range available {
		for _, nItem := range needed {
			if aItem.Quantity < nItem.Quantity {
				return errors.New("not enough item present")
			}
		}
	}
	return nil
}

func (os *OrderManagementService) checkDishAvailability(dishName string, menuType MenuType) (*Dish, error) {
	menu, err := os.restaurant.GetMenuByType(menuType)
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	dish, err := menu.GetDishByName(dishName)
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	if dish.Availability == DISH_NOT_AVAILABLE {
		return nil, errors.New("right now dish is not available")
	}

	neededItems, neededItemsNames, err := dish.GetItemsWithNames()
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	availableItems, notAvailableItems := os.ims.GetItemsByName(neededItemsNames)
	if len(notAvailableItems) > 0 {
		var errs []error
		for _, item := range notAvailableItems {
			errs = append(errs, fmt.Errorf("item: %s", item))
		}
		err := errors.Join(errs...)
		return nil, fmt.Errorf("CheckDishAvailability: dish will not be available: %w items are not available", err)
	}

	err = checkQuantity(availableItems, neededItems)
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: dish is not available: %w", err)
	}

	return dish, nil
}

func (os *OrderManagementService) allAvailable(dishName []string, menuType MenuType) ([]*Dish, error) {

	var foundDish []*Dish
	var errs []error

	for _, dName := range dishName {
		dish, err := os.checkDishAvailability(dName, menuType)
		if err != nil {
			errs = append(errs, fmt.Errorf("dish %s: %w", dName, err))
		} else {
			foundDish = append(foundDish, dish)
		}
	}

	if len(errs) > 0 {
		err := errors.Join(errs...)
		return foundDish, err
	}
	return foundDish, nil
}

func (os *OrderManagementService) OrderFood(dishName []string, menuType MenuType, paymentMode PaymentMode) (*Bill, error) {

	foundDish, err := os.allAvailable(dishName, menuType)
	if err != nil {
		return nil, fmt.Errorf("OrderFood: %w", err)
	}

	order := NewOrder(GenerateUUID(), foundDish)

	os.orders = append(os.orders, order)

	paymentService, err := os.paymentFactory.GetPaymentService(paymentMode)
	if err != nil {
		return nil, fmt.Errorf("PaymentMode %s: %w", paymentMode, err)
	}

	bill, err := paymentService.Payment(order)
	if err != nil {
		return nil, fmt.Errorf("payment failed!! try again: %w", err)
	}

	order.UpdateStatus(Processed)
	return bill, nil
}
