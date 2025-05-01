package order

import (
	"DesignRestaurantManagementSystem/factory"
	"DesignRestaurantManagementSystem/interfaces"
	billModel "DesignRestaurantManagementSystem/models/bill"
	dishModel "DesignRestaurantManagementSystem/models/dish"
	itemModel "DesignRestaurantManagementSystem/models/item"
	menuModel "DesignRestaurantManagementSystem/models/menu"
	orderModel "DesignRestaurantManagementSystem/models/order"
	paymentModel "DesignRestaurantManagementSystem/models/payment"
	"DesignRestaurantManagementSystem/utils"
	"errors"
	"fmt"
)

type OrderManagementService struct {
	paymentFactory *factory.PaymentFactory
	inventory      interfaces.InventoryManagementServiceInterface
	restaurant     interfaces.RestaurantServiceInterface
	orders         []*orderModel.Order
}

func NewOrderManagementService(paymentFactory *factory.PaymentFactory, inventory interfaces.InventoryManagementServiceInterface, restaurant interfaces.RestaurantServiceInterface) interfaces.OrderManagementServiceInterface {
	return &OrderManagementService{
		paymentFactory: paymentFactory,
		inventory:      inventory,
		restaurant:     restaurant,
		orders:         make([]*orderModel.Order, 0),
	}
}

func checkQuantity(available []*itemModel.Item, needed []*itemModel.Item) error {

	for _, aItem := range available {
		for _, nItem := range needed {
			if aItem.Quantity < nItem.Quantity {
				return errors.New("not enough item present")
			}
		}
	}
	return nil
}

func (os *OrderManagementService) checkDishAvailability(dishName string, menuType menuModel.MenuType) (*dishModel.Dish, error) {
	menu, err := os.restaurant.GetMenuByType(menuType)
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	dish, err := menu.GetDishByName(dishName)
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	if dish.Availability == dishModel.DISH_NOT_AVAILABLE {
		return nil, errors.New("right now dish is not available")
	}

	neededItems, neededItemsNames, err := dish.GetItemsWithNames()
	if err != nil {
		return nil, fmt.Errorf("CheckDishAvailability: %w", err)
	}

	availableItems, notAvailableItems := os.inventory.GetItemsByName(neededItemsNames)
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

func (os *OrderManagementService) allAvailable(dishName []string, menuType menuModel.MenuType) ([]*dishModel.Dish, error) {

	var foundDish []*dishModel.Dish
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

func (os *OrderManagementService) PlaceOrder(dishName []string, menuType menuModel.MenuType, paymentMode paymentModel.PaymentMode) (*billModel.Bill, error) {

	foundDish, err := os.allAvailable(dishName, menuType)
	if err != nil {
		return nil, fmt.Errorf("PlaceOrder: %w", err)
	}

	order := orderModel.NewOrder(utils.GenerateUUID(), foundDish)

	os.orders = append(os.orders, order)

	paymentService, err := os.paymentFactory.GetPaymentService(paymentMode)
	if err != nil {
		return nil, fmt.Errorf("PaymentMode %s: %w", paymentMode, err)
	}

	bill, err := paymentService.ProcessPayment(order)
	if err != nil {
		return nil, fmt.Errorf("payment failed!! try again: %w", err)
	}

	order.UpdateStatus(orderModel.Processed)
	return bill, nil
}

func (os *OrderManagementService) GetOrderStatus(orderID string) (orderModel.OrderStatus, error) {

	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return orderModel.OnProcess, fmt.Errorf("GetOrderStatus: %w", err)
	}

	return order.Status, nil
}

func (os *OrderManagementService) CancelOrder(orderID string) error {
	order, err := os.GetOrderByID(orderID)
	if err != nil {
		return fmt.Errorf("CancelOrder: %w", err)
	}

	if order.Status == orderModel.Processed {
		return errors.New("cannot cancel processed order")
	}

	order.UpdateStatus(orderModel.Canceled)
	return nil
}

func (os *OrderManagementService) GetOrderHistory() ([]*orderModel.Order, error) {
	if len(os.orders) == 0 {
		return nil, errors.New("no orders found")
	}
	return os.orders, nil
}

func (os *OrderManagementService) GetOrderByID(orderID string) (*orderModel.Order, error) {
	for _, order := range os.orders {
		if order.OrderID == orderID {
			return order, nil
		}
	}
	return nil, errors.New("order not found")
}
