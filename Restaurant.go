package main

import (
	"errors"
	"fmt"
)

type Restaurant struct {
	Name  string
	Menu  []*Menu
	Table []*Table
}

func NewRestaurant(name string) *Restaurant {
	return &Restaurant{
		Name:  name,
		Menu:  make([]*Menu, 0),
		Table: make([]*Table, 0),
	}
}

func (r *Restaurant) AddMenu(menu *Menu) {
	r.Menu = append(r.Menu, menu)
}

func (r *Restaurant) GetMenuByType(menuType MenuType) (*Menu, error) {
	for _, menu := range r.Menu {
		if menu.Type == menuType {
			return menu, nil
		}
	}
	return nil, errors.New("no menu found")
}

func (r *Restaurant) AddTable(table *Table) {
	r.Table = append(r.Table, table)
}

func (r *Restaurant) GetTableByID(tableID string) (*Table, error) {
	for _, table := range r.Table {
		if table.TableID == tableID {
			return table, nil
		}
	}
	return nil, errors.New("no table found")
}

func (r *Restaurant) ReserveTable(tableIDs []string) error {

	var tables []*Table
	for _, tableID := range tableIDs {
		table, err := r.GetTableByID(tableID)
		if err != nil {
			return fmt.Errorf("table %s: %w", tableID, err)
		}
		tables = append(tables, table)
	}

	for _, table := range tables {
		if table.Status == BOOKED {
			return fmt.Errorf("table %s is not available", table.TableID)
		}
		table.Status = BOOKED
	}
	return nil
}
