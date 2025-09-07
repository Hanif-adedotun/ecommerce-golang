To run the tests:
1. Create a test database connection
2. Set up your test environment
3. Run `go test -v` in the terminal
4. The tests will run and report any failures or errors

## Required Packages
- testing
- database/sql
- encoding/json
- math/rand
- net/http
- time
- github.com/google/uuid
- github.com/hanif-adedotun/ecommerce-golang/db
- github.com/hanif-adedotun/ecommerce-golang/model
```go
// go test -v ./...
```
To run the tests, navigate to the root directory of the project and run the command `go test -v ./...`. This will execute all the test files in the project.

## Required Packages
- github.com/hanif-adedotun/ecommerce-golang/db
- net/http
- database/sql
- context
- fmt
- log
- time
- testing
# TestLoadRoutes
## Description
This test function is used to test the `loadRoutes` function in the `application` package.
It checks if the status code of the response is 200 when a GET request is made to the "/" endpoint.
## TestLoadOrderRoute
## Description
This test function is used to test the `loadOrderRoute` function in the `application` package.
It checks if the status code of the response is 200 when a GET request is made to the "/orders" endpoint.

## Required Packages
- net/http
- net/http/httptest
- testing
- github.com/hanif-adedotun/ecommerce-golang/db
- github.com/hanif-adedotun/ecommerce-golang/handler
- github.com/go-chi/chi/v5
- github.com/go-chi/chi/v5/middleware
This code contains three test cases for the ConnectDB function: 
1. A successful database connection. 
2. Missing environment variables. 
3. An invalid database URI. 
These test cases cover the main scenarios that the ConnectDB function may encounter.

## Required Packages
- database/sql
- errors
- os
- testing
### Test Function: `TestPostgreRepo_Insert`

*   **Purpose:** Test the `Insert` method of the `PostgreRepo` struct.
*   **Input:** A `model.Order` object and a boolean indicating whether an error is expected.
*   **Output:** An error value.
*   **Test Cases:**
    *   Valid order: Test that a valid order can be inserted successfully.
    *   Invalid order: Test that an invalid order returns an error.


## Required Packages
- database/sql
- encoding/json
- github.com/hanif-adedotun/ecommerce-golang/model
- testing
### Test Cases for `app.go`

The following test cases are generated for the `app.go` file:

#### TestAppStart
*   Test that the `Start` method of the `App` struct does not return an error when called with a valid context.
*   Test that the `Start` method of the `App` struct returns an error when called with an invalid context.

#### TestLoadRoutes
*   Test that the `loadRoutes` method of the `App` struct sets the `router` field to a non-nil value.


## Required Packages
- net/http
- github.com/go-chi/chi/v5
- github.com/go-chi/chi/v5/middleware
- github.com/hanif-adedotun/ecommerce-golang/db
- github.com/hanif-adedotun/ecommerce-golang/handler
### TestConnectDB Function

This test function checks the ConnectDB function in the db package. It tests the function with both valid and invalid database connections.

#### Parameters

* `t *testing.T`: The test object.

#### Test Scenarios

1. Valid database connection: The test connects to a valid database and checks that no error is returned.
2. Invalid database connection: The test tries to connect to an invalid database and checks that an error is returned.

#### Assertions

The test checks that the error returned by the ConnectDB function matches the expected error.


## Required Packages
- database/sql
- testing
### Unit Tests for Order and LineItem structs

The following tests are designed to verify the correctness of the Order and LineItem structs.

#### TestOrder function tests the Order struct
- It creates a new Order with a specified OrderID, CustomerID, LineItems, and CreatedAt.
- Then it checks if the OrderID, CustomerID, LineItems, and CreatedAt are set correctly.

#### TestLineItem function tests the LineItem struct
- It creates a new LineItem with a specified ItemID, Quantity, and Price.
- Then it checks if the ItemID, Quantity, and Price are set correctly.


## Required Packages
- reflect
- testing
- time
- github.com/google/uuid
