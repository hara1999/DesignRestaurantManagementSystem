#!/bin/bash

# Script to run tests for the Restaurant Management System

# Set default values
VERBOSE=false
SPECIFIC_TEST=""
PACKAGE_PATH="./..."
COVERAGE=false

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    -v|--verbose)
      VERBOSE=true
      shift
      ;;
    -t|--test)
      SPECIFIC_TEST=$2
      shift 2
      ;;
    -p|--package)
      PACKAGE_PATH=$2
      shift 2
      ;;
    -c|--coverage)
      COVERAGE=true
      shift
      ;;
    -h|--help)
      echo "Usage: $0 [options]"
      echo "Options:"
      echo "  -v, --verbose          Enable verbose output"
      echo "  -t, --test TEST_NAME   Run a specific test function"
      echo "  -p, --package PACKAGE  Run tests in a specific package path"
      echo "  -c, --coverage         Generate code coverage report"
      echo "  -h, --help             Display this help message"
      echo
      echo "Examples:"
      echo "  $0 -v                            # Run all tests with verbose output"
      echo "  $0 -t TestAddMenu -p ./services/restaurant/...  # Run a specific test function"
      echo "  $0 -p ./handlers/...            # Run tests in a specific package"
      echo "  $0 -c                            # Generate code coverage report"
      exit 0
      ;;
    *)
      echo "Unknown option: $1"
      echo "Use '$0 --help' for usage information"
      exit 1
      ;;
  esac
done

# Construct the test command
TEST_CMD="go test"

if [ "$VERBOSE" = true ]; then
  TEST_CMD="$TEST_CMD -v"
fi

if [ -n "$SPECIFIC_TEST" ]; then
  TEST_CMD="$TEST_CMD -run $SPECIFIC_TEST"
fi

if [ "$COVERAGE" = true ]; then
  TEST_CMD="$TEST_CMD -coverprofile=coverage.out"
fi

TEST_CMD="$TEST_CMD $PACKAGE_PATH"

# Display and run the command
echo "Running: $TEST_CMD"
eval $TEST_CMD

# Display coverage report if requested
if [ "$COVERAGE" = true ]; then
  echo "Generating coverage report..."
  go tool cover -html=coverage.out -o coverage.html
  echo "Coverage report saved as coverage.html"
  
  # Try to open the coverage report in a browser
  case "$(uname -s)" in
    Darwin)
      open coverage.html
      ;;
    Linux)
      if command -v xdg-open > /dev/null; then
        xdg-open coverage.html
      else
        echo "Please open coverage.html in your browser to view the coverage report"
      fi
      ;;
    *)
      echo "Please open coverage.html in your browser to view the coverage report"
      ;;
  esac
fi 