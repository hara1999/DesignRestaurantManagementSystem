# Restaurant Management System Tests

This directory contains unit tests for the Restaurant Management System.

## Running the Tests

To run all tests in the project, use the following command from the project root:

```bash
go test ./...
```

To run tests for a specific package, use:

```bash
go test ./handlers/...  # Run all handler tests
go test ./services/restaurant/...  # Run restaurant service tests
```

To run a specific test function, use:

```bash
go test -run TestAddMenu ./services/restaurant/...
```

To see verbose output, add the `-v` flag:

```bash
go test -v ./...
```

## Test Structure

The tests are structured as follows:

### Mock Implementations

- `tests/mocks/base_handler_mock.go`: Mock implementation of the BaseHandlerInterface
- `tests/mocks/restaurant_service_mock.go`: Mock implementation of the RestaurantServiceInterface

### Handler Tests

- `handlers/restaurant_handler_test.go`: Tests for the RestaurantHandler
  - Tests all API endpoints defined in RestaurantHandlerInterface
  - Uses mock service and base handler implementations
  - Tests both success and error cases for each endpoint

### Service Tests

- `services/restaurant/restaurant_test.go`: Tests for the RestaurantService 
  - Tests all service methods defined in RestaurantServiceInterface
  - Tests both success and error cases for each method

## Adding New Tests

When adding new features:

1. Add unit tests for new service methods
2. Add unit tests for new handler endpoints
3. Run the tests to ensure everything works correctly 