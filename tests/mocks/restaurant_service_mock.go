package mocks

import (
	"DesignRestaurantManagementSystem/interfaces"
	dishModel "DesignRestaurantManagementSystem/models/dish"
	menuModel "DesignRestaurantManagementSystem/models/menu"
	tableModel "DesignRestaurantManagementSystem/models/table"
)

// MockRestaurantService is a mock implementation of RestaurantServiceInterface
type MockRestaurantService struct {
	// AddMenuFunc will be called when AddMenu is called
	AddMenuFunc func(menuReq menuModel.AddMenuRequest) error

	// AddTableFunc will be called when AddTable is called
	AddTableFunc func(tableReq tableModel.AddTableRequest) error

	// AddDishFunc will be called when AddDish is called
	AddDishFunc func(dishReq dishModel.AddDishRequest) error

	// GetMenuByTypeFunc will be called when GetMenuByType is called
	GetMenuByTypeFunc func(menuType string) (*menuModel.Menu, error)

	// GetTableByIDFunc will be called when GetTableByID is called
	GetTableByIDFunc func(tableID string) (*tableModel.Table, error)

	// GetTablesFunc will be called when GetTables is called
	GetTablesFunc func() ([]*tableModel.Table, error)

	// ReserveTableFunc will be called when ReserveTable is called
	ReserveTableFunc func(tableIDs []string) error

	// UpdateDishFunc will be called when UpdateDish is called
	UpdateDishFunc func(dishReq dishModel.AddDishRequest) error

	// RemoveDishFunc will be called when RemoveDish is called
	RemoveDishFunc func(menuType string, dishName string) error
}

// Ensure MockRestaurantService implements RestaurantServiceInterface
var _ interfaces.RestaurantServiceInterface = (*MockRestaurantService)(nil)

func (m *MockRestaurantService) AddMenu(menuReq menuModel.AddMenuRequest) error {
	if m.AddMenuFunc != nil {
		return m.AddMenuFunc(menuReq)
	}
	return nil
}

func (m *MockRestaurantService) AddTable(tableReq tableModel.AddTableRequest) error {
	if m.AddTableFunc != nil {
		return m.AddTableFunc(tableReq)
	}
	return nil
}

func (m *MockRestaurantService) AddDish(dishReq dishModel.AddDishRequest) error {
	if m.AddDishFunc != nil {
		return m.AddDishFunc(dishReq)
	}
	return nil
}

func (m *MockRestaurantService) GetMenuByType(menuType string) (*menuModel.Menu, error) {
	if m.GetMenuByTypeFunc != nil {
		return m.GetMenuByTypeFunc(menuType)
	}
	return nil, nil
}

func (m *MockRestaurantService) GetTableByID(tableID string) (*tableModel.Table, error) {
	if m.GetTableByIDFunc != nil {
		return m.GetTableByIDFunc(tableID)
	}
	return nil, nil
}

func (m *MockRestaurantService) GetTables() ([]*tableModel.Table, error) {
	if m.GetTablesFunc != nil {
		return m.GetTablesFunc()
	}
	return nil, nil
}

func (m *MockRestaurantService) ReserveTable(tableIDs []string) error {
	if m.ReserveTableFunc != nil {
		return m.ReserveTableFunc(tableIDs)
	}
	return nil
}

func (m *MockRestaurantService) UpdateDish(dishReq dishModel.AddDishRequest) error {
	if m.UpdateDishFunc != nil {
		return m.UpdateDishFunc(dishReq)
	}
	return nil
}

func (m *MockRestaurantService) RemoveDish(menuType string, dishName string) error {
	if m.RemoveDishFunc != nil {
		return m.RemoveDishFunc(menuType, dishName)
	}
	return nil
}
