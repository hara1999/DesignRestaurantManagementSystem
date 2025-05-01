package restaurant

import (
	"errors"
	"fmt"

	menuModel "DesignRestaurantManagementSystem/models/menu"
	tableModel "DesignRestaurantManagementSystem/models/table"
)

type RestaurantService struct {
	Name  string
	Menu  []*menuModel.Menu
	Table []*tableModel.Table
}

func NewRestaurant(name string) *RestaurantService {
	return &RestaurantService{
		Name:  name,
		Menu:  make([]*menuModel.Menu, 0),
		Table: make([]*tableModel.Table, 0),
	}
}

func (r *RestaurantService) AddMenu(menu *menuModel.Menu) {
	r.Menu = append(r.Menu, menu)
}

func (r *RestaurantService) GetMenuByType(menuType menuModel.MenuType) (*menuModel.Menu, error) {
	for _, menu := range r.Menu {
		if menu.Type == menuType {
			return menu, nil
		}
	}
	return nil, errors.New("no menu found")
}

func (r *RestaurantService) AddTable(table *tableModel.Table) {

	r.Table = append(r.Table, table)
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
