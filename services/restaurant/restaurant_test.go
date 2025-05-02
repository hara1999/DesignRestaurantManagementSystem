package restaurant

import (
	"DesignRestaurantManagementSystem/models/dish"
	"DesignRestaurantManagementSystem/models/item"
	"DesignRestaurantManagementSystem/models/menu"
	"DesignRestaurantManagementSystem/models/table"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMenu(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Test case: Add a new menu
	menuReq := menu.AddMenuRequest{
		MenuType: "MainCourse",
		Dish: []dish.AddDishRequest{
			{
				Name:     "Pizza",
				Price:    200,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Cheese", Quantity: 2},
					{Name: "Tomato", Quantity: 1},
				},
			},
		},
	}

	err := restaurantService.AddMenu(menuReq)
	assert.NoError(t, err)

	// Verify the menu was added
	menuFromService, err := restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	assert.Equal(t, menu.MenuType("MainCourse"), menuFromService.Type)
	assert.Len(t, menuFromService.Dish, 1)
	assert.Equal(t, "Pizza", menuFromService.Dish[0].DishName)
}

func TestAddDish(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// First, add a menu
	menuReq := menu.AddMenuRequest{
		MenuType: "MainCourse",
		Dish:     []dish.AddDishRequest{},
	}
	err := restaurantService.AddMenu(menuReq)
	assert.NoError(t, err)

	// Test case: Add a dish to existing menu
	dishReq := dish.AddDishRequest{
		Name:     "Pasta",
		Price:    150,
		MenuType: "MainCourse",
		Items: []item.AddItemRequest{
			{Name: "Pasta", Quantity: 1},
			{Name: "Sauce", Quantity: 1},
		},
	}

	err = restaurantService.AddDish(dishReq)
	assert.NoError(t, err)

	// Verify the dish was added
	menuFromService, err := restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	assert.Len(t, menuFromService.Dish, 1)
	assert.Equal(t, "Pasta", menuFromService.Dish[0].DishName)
	assert.Equal(t, 150, menuFromService.Dish[0].Price)
	assert.Len(t, menuFromService.Dish[0].Item, 2)

	// Test case: Add a dish to non-existent menu
	invalidDishReq := dish.AddDishRequest{
		Name:     "InvalidDish",
		Price:    100,
		MenuType: "NonExistentMenu",
		Items: []item.AddItemRequest{
			{Name: "Item1", Quantity: 1},
		},
	}

	err = restaurantService.AddDish(invalidDishReq)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no menu found")
}

func TestAddTable(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Test case: Add a table
	tableReq := table.AddTableRequest{
		TableID: "T001",
	}

	err := restaurantService.AddTable(tableReq)
	assert.NoError(t, err)

	// Verify the table was added
	tableFromService, err := restaurantService.GetTableByID("T001")
	assert.NoError(t, err)
	assert.Equal(t, "T001", tableFromService.TableID)
	assert.Equal(t, table.AVAILABLE, tableFromService.Status)

	// Test case: Get a non-existent table
	_, err = restaurantService.GetTableByID("NonExistentTable")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no table found")
}

func TestGetMenuByType(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Add a menu
	menuReq := menu.AddMenuRequest{
		MenuType: "MainCourse",
		Dish: []dish.AddDishRequest{
			{
				Name:     "Pizza",
				Price:    200,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Cheese", Quantity: 2},
				},
			},
		},
	}
	err := restaurantService.AddMenu(menuReq)
	assert.NoError(t, err)

	// Test case: Get an existing menu
	menuFromService, err := restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	assert.Equal(t, menu.MenuType("MainCourse"), menuFromService.Type)
	assert.Len(t, menuFromService.Dish, 1)

	// Test case: Get a non-existent menu
	_, err = restaurantService.GetMenuByType("NonExistentMenu")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no menu found")
}

func TestGetTables(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Test case: Get tables when none exist
	tables, err := restaurantService.GetTables()
	assert.Error(t, err)
	assert.Nil(t, tables)
	assert.Contains(t, err.Error(), "no tables found")

	// Add tables
	tableReq1 := table.AddTableRequest{TableID: "T001"}
	tableReq2 := table.AddTableRequest{TableID: "T002"}
	err = restaurantService.AddTable(tableReq1)
	assert.NoError(t, err)
	err = restaurantService.AddTable(tableReq2)
	assert.NoError(t, err)

	// Test case: Get tables after adding
	tables, err = restaurantService.GetTables()
	assert.NoError(t, err)
	assert.Len(t, tables, 2)
	assert.Equal(t, "T001", tables[0].TableID)
	assert.Equal(t, "T002", tables[1].TableID)
}

func TestReserveTable(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Add tables
	tableReq1 := table.AddTableRequest{TableID: "T001"}
	tableReq2 := table.AddTableRequest{TableID: "T002"}
	err := restaurantService.AddTable(tableReq1)
	assert.NoError(t, err)
	err = restaurantService.AddTable(tableReq2)
	assert.NoError(t, err)

	// Test case: Reserve existing tables
	err = restaurantService.ReserveTable([]string{"T001", "T002"})
	assert.NoError(t, err)

	// Verify the tables were reserved
	table1, err := restaurantService.GetTableByID("T001")
	assert.NoError(t, err)
	assert.Equal(t, table.BOOKED, table1.Status)

	table2, err := restaurantService.GetTableByID("T002")
	assert.NoError(t, err)
	assert.Equal(t, table.BOOKED, table2.Status)

	// Test case: Reserve a non-existent table
	err = restaurantService.ReserveTable([]string{"T003"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no table found")

	// Test case: Reserve an already booked table
	err = restaurantService.ReserveTable([]string{"T001"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "is not available")
}

func TestUpdateDish(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Add a menu with a dish
	menuReq := menu.AddMenuRequest{
		MenuType: "MainCourse",
		Dish: []dish.AddDishRequest{
			{
				Name:     "Pizza",
				Price:    200,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Cheese", Quantity: 2},
					{Name: "Tomato", Quantity: 1},
				},
			},
		},
	}
	err := restaurantService.AddMenu(menuReq)
	assert.NoError(t, err)

	// Test case: Update an existing dish
	updateDishReq := dish.AddDishRequest{
		Name:     "Pizza",
		Price:    250, // Updated price
		MenuType: "MainCourse",
		Items: []item.AddItemRequest{
			{Name: "Cheese", Quantity: 3}, // Updated quantity
			{Name: "Tomato", Quantity: 2},
			{Name: "Basil", Quantity: 1}, // Added new item
		},
	}

	err = restaurantService.UpdateDish(updateDishReq)
	assert.NoError(t, err)

	// Verify the dish was updated
	menuFromService, err := restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	updatedDish, err := menuFromService.GetDishByName("Pizza")
	assert.NoError(t, err)
	assert.Equal(t, 250, updatedDish.Price)
	assert.Len(t, updatedDish.Item, 3)

	// Test case: Update a non-existent dish
	nonExistentDishReq := dish.AddDishRequest{
		Name:     "NonExistentDish",
		Price:    150,
		MenuType: "MainCourse",
		Items:    []item.AddItemRequest{},
	}

	err = restaurantService.UpdateDish(nonExistentDishReq)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no dish found")

	// Test case: Update a dish in a non-existent menu
	invalidMenuDishReq := dish.AddDishRequest{
		Name:     "Pizza",
		Price:    300,
		MenuType: "NonExistentMenu",
		Items:    []item.AddItemRequest{},
	}

	err = restaurantService.UpdateDish(invalidMenuDishReq)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no menu found")
}

func TestRemoveDish(t *testing.T) {
	restaurantService := NewRestaurantService("Test Restaurant")

	// Add a menu with multiple dishes
	menuReq := menu.AddMenuRequest{
		MenuType: "MainCourse",
		Dish: []dish.AddDishRequest{
			{
				Name:     "Pizza",
				Price:    200,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Cheese", Quantity: 2},
				},
			},
			{
				Name:     "Pasta",
				Price:    150,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Pasta", Quantity: 1},
				},
			},
		},
	}
	err := restaurantService.AddMenu(menuReq)
	assert.NoError(t, err)

	// Verify both dishes exist
	menuFromService, err := restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	assert.Len(t, menuFromService.Dish, 2)

	// Test case: Remove an existing dish
	err = restaurantService.RemoveDish("MainCourse", "Pizza")
	assert.NoError(t, err)

	// Verify the dish was removed
	menuFromService, err = restaurantService.GetMenuByType("MainCourse")
	assert.NoError(t, err)
	assert.Len(t, menuFromService.Dish, 1)
	assert.Equal(t, "Pasta", menuFromService.Dish[0].DishName)

	// Test case: Remove a non-existent dish
	err = restaurantService.RemoveDish("MainCourse", "NonExistentDish")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no dish found")

	// Test case: Remove a dish from a non-existent menu
	err = restaurantService.RemoveDish("NonExistentMenu", "Pasta")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no menu found")
}
