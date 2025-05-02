# DesignRestaurantManagementSystem
LLD design for Restaurant Management System


This **Low-Level Design (LLD)** enables customers to **view the menu**, **place food orders**, and **make table reservations** in advance.

The system includes robust **order processing**, **billing**, and **payment handling** functionalities. While the current implementation supports a **single payment method**, the design is easily extendable to support **multiple payment options** in the future.

An `InventoryManagementService` is responsible for efficiently managing **inventory operations**.

Additionally, the system manages **staff information**, including their **roles** and **schedules**.

## Testing

This project includes a comprehensive test suite to ensure all functionality works as expected:

- **Unit Tests**: Tests for individual components (services, handlers)
- **Test Coverage**: Script to generate a coverage report

### Running Tests

You can run the tests using the provided script:

```bash
# Make the script executable (if not already)
chmod +x scripts/run_tests.sh

# Run all tests
./scripts/run_tests.sh

# Run tests with verbose output
./scripts/run_tests.sh -v

# Run specific test
./scripts/run_tests.sh -t TestAddMenu -p ./services/restaurant/...

# Generate coverage report
./scripts/run_tests.sh -c
```

For more options, run:

```bash
./scripts/run_tests.sh -h
```

See `tests/README.md` for more information about the test structure and how to add new tests.


