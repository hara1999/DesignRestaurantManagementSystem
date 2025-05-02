package restaurant

import (
	"errors"
	"fmt"

	"DesignRestaurantManagementSystem/interfaces"
	dishModel "DesignRestaurantManagementSystem/models/dish"
	itemModel "DesignRestaurantManagementSystem/models/item"
	menuModel "DesignRestaurantManagementSystem/models/menu"
	tableModel "DesignRestaurantManagementSystem/models/table"
)

type RestaurantService struct {
	Name  string
	Menu  []*menuModel.Menu
	Table []*tableModel.Table
}

func NewRestaurantService(name string) interfaces.RestaurantServiceInterface {
	return &RestaurantService{
		Name:  name,
		Menu:  make([]*menuModel.Menu, 0),
		Table: make([]*tableModel.Table, 0),
	}
}

func (r *RestaurantService) AddMenu(menuReq menuModel.AddMenuRequest) error {

	for _, menu := range r.Menu {
		if menu.Type == menuModel.MenuType(menuReq.MenuType) {
			return fmt.Errorf("menu %s already exists", menuReq.MenuType)
		}
	}

	menu := menuModel.NewMenu(menuModel.MenuType(menuReq.MenuType))
	r.Menu = append(r.Menu, menu)
	for _, dish := range menuReq.Dish {
		err := r.AddDish(dish)
		if err != nil {
			return fmt.Errorf("dish %s: %w", dish.Name, err)
		}
	}
	return nil
}

func (r *RestaurantService) AddDish(dishReq dishModel.AddDishRequest) error {
	menu, err := r.GetMenuByType(dishReq.MenuType)
	if err != nil {
		return fmt.Errorf("menu %s: %w", dishReq.MenuType, err)
	}

	for _, dish := range menu.Dish {
		if dish.DishName == dishReq.Name {
			return fmt.Errorf("dish %s already exists", dishReq.Name)
		}
	}

	var items []*itemModel.Item
	for _, item := range dishReq.Items {
		items = append(items, itemModel.NewItem(item.Name, item.Quantity))
	}

	dish := dishModel.NewDish(dishReq.Name, dishReq.Price, items)
	menu.AddDish(dish)
	return nil
}

func (r *RestaurantService) AddTable(tableReq tableModel.AddTableRequest) error {
	table := tableModel.NewTable(tableReq.TableID)
	r.Table = append(r.Table, table)
	return nil
}

func (r *RestaurantService) GetMenuByType(menuType string) (*menuModel.Menu, error) {
	for _, menu := range r.Menu {
		if menu.Type == menuModel.MenuType(menuType) {
			return menu, nil
		}
	}
	return nil, errors.New("no menu found")
}

func (r *RestaurantService) GetTableByID(tableID string) (*tableModel.Table, error) {
	for _, table := range r.Table {
		if table.TableID == tableID {
			return table, nil
		}
	}
	return nil, errors.New("no table found")
}

func (r *RestaurantService) GetTables() ([]*tableModel.Table, error) {
	if len(r.Table) == 0 {
		return nil, errors.New("no tables found")
	}
	return r.Table, nil
}

func (r *RestaurantService) ReserveTable(tableIDs []string) error {

	var tables []*tableModel.Table
	for _, tableID := range tableIDs {
		table, err := r.GetTableByID(tableID)
		if err != nil {
			return fmt.Errorf("table %s: %w", tableID, err)
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		if table.Status == tableModel.BOOKED {
			return fmt.Errorf("table %s is not available", table.TableID)
		}
		table.Status = tableModel.BOOKED
	}
	return nil
}

func (r *RestaurantService) UpdateDish(dishReq dishModel.AddDishRequest) error {
	menu, err := r.GetMenuByType(dishReq.MenuType)
	if err != nil {
		return fmt.Errorf("menu %s: %w", dishReq.MenuType, err)
	}

	dish, err := menu.GetDishByName(dishReq.Name)
	if err != nil {
		return fmt.Errorf("dish %s: %w", dishReq.Name, err)
	}

	dish.Price = dishReq.Price

	if len(dishReq.Items) > 0 {
		dish.Item = make([]*itemModel.Item, 0)
		for _, item := range dishReq.Items {
			dish.AddItem(itemModel.NewItem(item.Name, item.Quantity))
		}
	}

	return nil
}

func (r *RestaurantService) RemoveDish(menuType string, dishName string) error {
	menu, err := r.GetMenuByType(menuType)
	if err != nil {
		return fmt.Errorf("menu %s: %w", menuType, err)
	}

	dish, err := menu.GetDishByName(dishName)
	if err != nil {
		return fmt.Errorf("dish %s: %w", dishName, err)
	}

	menu.RemoveDish(dish)
	return nil
}
