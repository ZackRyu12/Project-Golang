# Enigma Laundry Management System

Enigma Laundry Management System is a simple CLI-based application written in Go to manage customers, services, and orders in a laundry business.

## Features

- **Customer Management**: Create, view, update, and delete customer records.
- **Service Management**: Create, view, update, and delete laundry services.
- **Order Management**: Create, view, complete, and manage orders.

## Requirements

- PostgreSQL installed and running
- Go 1.16+ installed

## Database Setup

1. Install PostgreSQL and ensure it's running.
2. Create a database named `enigma_laundry`.
3. Import the DDL file (`DDL.sql`) to set up the database schema.
4. Import the DML file (DML.sql) to populate initial data.

   ```bash
   psql -U postgres -d enigma_laundry -f DDL.sql

  ## How to Run the Application
1. Clone the repository or download the project files.
2. Ensure you have Go installed on your machine.
3. Navigate to the project directory and update the PostgreSQL credentials in the main.go file:

    ```bash
    const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "1234"
    dbname   = "enigma_laundry"
    )
4. Run the project using the following command:

    ```bash
    go run main.go

## Application Usage
When you run the application, the following menu will appear:

### Main Menu:

    ```bash
    Main Menu:
    1. Customer
    2. Service
    3. Order
    4. Exit

### Customer Menu:
You can manage customer information using this menu. Options include:
- Create Customer
- View Customer List
- View Customer Details
- Update Customer
- Delete Customer

### Service Menu:
You can manage laundry services using this menu. Options include:

- Create Service
- View Service List
- View Service Details
- Update Service
- Delete Service
### Order Menu:
You can manage customer orders using this menu. Options include:

- Create Order
- Complete Order
- View Order List
- View Order Details
### Error Handling
- If you try to delete a customer or service that is currently being used in an order, you will receive an error message:
    - For customers: Customer ID is being used in orders. Please delete the order first.
    - For services: Service ID is being used in orders. Please delete the order first.