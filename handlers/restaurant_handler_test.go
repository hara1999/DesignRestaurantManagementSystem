package handlers_test

import (
	"DesignRestaurantManagementSystem/handlers"
	"DesignRestaurantManagementSystem/interfaces"
	"DesignRestaurantManagementSystem/models/dish"
	"DesignRestaurantManagementSystem/models/item"
	"DesignRestaurantManagementSystem/models/menu"
	"DesignRestaurantManagementSystem/models/table"
	"DesignRestaurantManagementSystem/tests/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRestaurantHandler() (*mocks.MockBaseHandler, *mocks.MockRestaurantService, interfaces.RestaurantHandlerInterface, *gin.Engine) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create mocks
	mockBase := &mocks.MockBaseHandler{}
	mockService := &mocks.MockRestaurantService{}

	// Create handler with mocks
	handler := handlers.NewRestaurantHandler(mockBase, mockService)

	// Setup router
	router := gin.Default()

	return mockBase, mockService, handler, router
}

func TestAddMenu(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/menu/add", handler.AddMenu)

	tests := []struct {
		name           string
		requestBody    menu.AddMenuRequest
		serviceError   error
		expectedStatus int
		expectedData   interface{}
	}{
		{
			name: "Success",
			requestBody: menu.AddMenuRequest{
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
			},
			serviceError:   nil,
			expectedStatus: 201,
		},
		{
			name: "Service Error",
			requestBody: menu.AddMenuRequest{
				MenuType: "InvalidMenu",
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
			},
			serviceError:   errors.New("invalid menu type"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.AddMenuFunc = func(req menu.AddMenuRequest) error {
				return tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/restaurant/menu/add", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestAddDish(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/dish/add", handler.AddDish)

	tests := []struct {
		name           string
		requestBody    dish.AddDishRequest
		serviceError   error
		expectedStatus int
	}{
		{
			name: "Success",
			requestBody: dish.AddDishRequest{
				Name:     "Pasta",
				Price:    150,
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Pasta", Quantity: 1},
					{Name: "Sauce", Quantity: 1},
				},
			},
			serviceError:   nil,
			expectedStatus: 201,
		},
		{
			name: "Service Error",
			requestBody: dish.AddDishRequest{
				Name:     "Pasta",
				Price:    150,
				MenuType: "InvalidMenu",
				Items: []item.AddItemRequest{
					{Name: "Pasta", Quantity: 1},
				},
			},
			serviceError:   errors.New("menu not found"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.AddDishFunc = func(req dish.AddDishRequest) error {
				return tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/restaurant/dish/add", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestAddTable(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/tables/add", handler.AddTable)

	tests := []struct {
		name           string
		requestBody    table.AddTableRequest
		serviceError   error
		expectedStatus int
	}{
		{
			name: "Success",
			requestBody: table.AddTableRequest{
				TableID: "T001",
			},
			serviceError:   nil,
			expectedStatus: 201,
		},
		{
			name: "Service Error",
			requestBody: table.AddTableRequest{
				TableID: "",
			},
			serviceError:   errors.New("invalid table ID"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.AddTableFunc = func(req table.AddTableRequest) error {
				return tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/restaurant/tables/add", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestGetTables(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.GET("/restaurant/tables", handler.GetTables)

	tests := []struct {
		name           string
		serviceTables  []*table.Table
		serviceError   error
		expectedStatus int
	}{
		{
			name: "Success",
			serviceTables: []*table.Table{
				{TableID: "T001", Status: table.AVAILABLE},
				{TableID: "T002", Status: table.BOOKED},
			},
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name:           "Service Error",
			serviceTables:  nil,
			serviceError:   errors.New("no tables found"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.GetTablesFunc = func() ([]*table.Table, error) {
				return tc.serviceTables, tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				if tc.serviceError == nil {
					assert.Equal(t, tc.serviceTables, data)
				}
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			req, _ := http.NewRequest(http.MethodGet, "/restaurant/tables", nil)

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestReserveTable(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/tables/reserve", handler.ReserveTable)

	tests := []struct {
		name           string
		requestBody    []string
		serviceError   error
		expectedStatus int
	}{
		{
			name:           "Success",
			requestBody:    []string{"T001", "T002"},
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name:           "Empty Table IDs",
			requestBody:    []string{},
			expectedStatus: 400,
		},
		{
			name:           "Service Error",
			requestBody:    []string{"T001", "T999"},
			serviceError:   errors.New("table T999 not found"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.ReserveTableFunc = func(tableIDs []string) error {
				if len(tableIDs) == 0 {
					return errors.New("empty table IDs")
				}
				return tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/restaurant/tables/reserve", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil && len(tc.requestBody) > 0 {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestGetMenuByType(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.GET("/restaurant/menu", handler.GetMenuByType)

	// Create sample menu for testing
	sampleMenu := &menu.Menu{
		Type: menu.MainCourse,
		Dish: []*dish.Dish{
			{DishName: "Pizza", Price: 200},
			{DishName: "Pasta", Price: 150},
		},
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		serviceMenu    *menu.Menu
		serviceError   error
		expectedStatus int
	}{
		{
			name:           "Success",
			queryParams:    map[string]string{"type": "MainCourse"},
			serviceMenu:    sampleMenu,
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name:           "Missing Type Parameter",
			queryParams:    map[string]string{},
			serviceMenu:    nil,
			serviceError:   nil,
			expectedStatus: 400,
		},
		{
			name:           "Menu Not Found",
			queryParams:    map[string]string{"type": "Dessert"},
			serviceMenu:    nil,
			serviceError:   errors.New("no menu found"),
			expectedStatus: 404,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.GetMenuByTypeFunc = func(menuType string) (*menu.Menu, error) {
				return tc.serviceMenu, tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				if tc.serviceError == nil && tc.serviceMenu != nil {
					assert.Equal(t, tc.serviceMenu, data)
				}
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqURL := "/restaurant/menu"
			if len(tc.queryParams) > 0 {
				reqURL += "?"
				for k, v := range tc.queryParams {
					reqURL += k + "=" + v + "&"
				}
				reqURL = reqURL[:len(reqURL)-1] // Remove the trailing &
			}

			req, _ := http.NewRequest(http.MethodGet, reqURL, nil)

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil && tc.serviceMenu != nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestUpdateDish(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/dish/update", handler.UpdateDish)

	tests := []struct {
		name           string
		requestBody    dish.AddDishRequest
		serviceError   error
		expectedStatus int
	}{
		{
			name: "Success",
			requestBody: dish.AddDishRequest{
				Name:     "Pizza",
				Price:    250, // Updated price
				MenuType: "MainCourse",
				Items: []item.AddItemRequest{
					{Name: "Cheese", Quantity: 3}, // Updated quantity
					{Name: "Tomato", Quantity: 2},
				},
			},
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name: "Service Error",
			requestBody: dish.AddDishRequest{
				Name:     "NonExistentDish",
				Price:    150,
				MenuType: "MainCourse",
				Items:    []item.AddItemRequest{},
			},
			serviceError:   errors.New("no dish found"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks for each test case
			mockBase.SuccessResponseFunc = nil
			mockBase.ErrorResponseFunc = nil
			mockBase.ValidationErrorFunc = nil

			// Setup mock service expectations
			mockService.UpdateDishFunc = func(req dish.AddDishRequest) error {
				return tc.serviceError
			}

			// Set up mock handler methods
			mockBase.ValidationErrorFunc = func(c *gin.Context, err error) {
				c.JSON(400, gin.H{
					"status":  "error",
					"message": "Validation failed",
					"errors":  err.Error(),
				})
			}

			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				c.JSON(status, gin.H{
					"status":  "error",
					"message": message,
				})
			}

			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				c.JSON(status, gin.H{
					"status": "success",
					"data":   data,
				})
			}

			// Prepare request
			reqBytes, _ := json.Marshal(tc.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/restaurant/dish/update", bytes.NewBuffer(reqBytes))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Check the response directly
			assert.Equal(t, tc.expectedStatus, w.Code, "HTTP status code should match expected")

			if tc.serviceError == nil {
				// For success case, check that the response contains success data
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "Should be able to unmarshal response")
				assert.Equal(t, "success", response["status"], "Response status should be 'success'")
			} else {
				// For error case, check that the response contains error data
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err, "Should be able to unmarshal response")
				assert.Equal(t, "error", response["status"], "Response status should be 'error'")
				assert.Contains(t, response["message"], tc.serviceError.Error(), "Error message should contain service error")
			}
		})
	}
}

func TestRemoveDish(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.POST("/restaurant/dish/remove", handler.RemoveDish)

	tests := []struct {
		name           string
		queryParams    map[string]string
		serviceError   error
		expectedStatus int
	}{
		{
			name:           "Success",
			queryParams:    map[string]string{"type": "MainCourse", "name": "Pizza"},
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name:           "Missing Parameters",
			queryParams:    map[string]string{"type": "MainCourse"}, // Missing name
			serviceError:   nil,
			expectedStatus: 400,
		},
		{
			name:           "Service Error",
			queryParams:    map[string]string{"type": "MainCourse", "name": "NonExistentDish"},
			serviceError:   errors.New("no dish found"),
			expectedStatus: 400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.RemoveDishFunc = func(menuType string, dishName string) error {
				return tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqURL := "/restaurant/dish/remove"
			if len(tc.queryParams) > 0 {
				reqURL += "?"
				for k, v := range tc.queryParams {
					reqURL += k + "=" + v + "&"
				}
				reqURL = reqURL[:len(reqURL)-1] // Remove the trailing &
			}

			req, _ := http.NewRequest(http.MethodPost, reqURL, nil)

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil && len(tc.queryParams) == 2 {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}

func TestGetTableByID(t *testing.T) {
	mockBase, mockService, handler, router := setupRestaurantHandler()

	// Setup route
	router.GET("/restaurant/table", handler.GetTableByID)

	// Create sample table for testing
	sampleTable := &table.Table{
		TableID: "T001",
		Status:  table.AVAILABLE,
	}

	tests := []struct {
		name           string
		queryParams    map[string]string
		serviceTable   *table.Table
		serviceError   error
		expectedStatus int
	}{
		{
			name:           "Success",
			queryParams:    map[string]string{"id": "T001"},
			serviceTable:   sampleTable,
			serviceError:   nil,
			expectedStatus: 200,
		},
		{
			name:           "Missing ID Parameter",
			queryParams:    map[string]string{},
			serviceTable:   nil,
			serviceError:   nil,
			expectedStatus: 400,
		},
		{
			name:           "Table Not Found",
			queryParams:    map[string]string{"id": "T999"},
			serviceTable:   nil,
			serviceError:   errors.New("no table found"),
			expectedStatus: 404,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock service expectations
			mockService.GetTableByIDFunc = func(tableID string) (*table.Table, error) {
				return tc.serviceTable, tc.serviceError
			}

			var successResponseCalled bool
			mockBase.SuccessResponseFunc = func(c *gin.Context, status int, data interface{}) {
				successResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				if tc.serviceError == nil && tc.serviceTable != nil {
					assert.Equal(t, tc.serviceTable, data)
				}
			}

			var errorResponseCalled bool
			mockBase.ErrorResponseFunc = func(c *gin.Context, status int, message string) {
				errorResponseCalled = true
				assert.Equal(t, tc.expectedStatus, status)
				assert.NotEmpty(t, message)
			}

			// Prepare request
			reqURL := "/restaurant/table"
			if len(tc.queryParams) > 0 {
				reqURL += "?"
				for k, v := range tc.queryParams {
					reqURL += k + "=" + v + "&"
				}
				reqURL = reqURL[:len(reqURL)-1] // Remove the trailing &
			}

			req, _ := http.NewRequest(http.MethodGet, reqURL, nil)

			// Perform request
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verify results
			if tc.serviceError == nil && tc.serviceTable != nil {
				assert.True(t, successResponseCalled, "SuccessResponse should be called")
			} else {
				assert.True(t, errorResponseCalled, "ErrorResponse should be called")
			}
		})
	}
}
