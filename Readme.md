CREATE TABLE customer (
    customer_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    address VARCHAR(255) DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE service (
    service_id INT PRIMARY KEY,
    service_name VARCHAR(255) NOT NULL,
    unit VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "order" (
    order_id INT PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date TIMESTAMP NOT NULL,
    completion_date TIMESTAMP,
    received_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(customer_id) REFERENCES customer(customer_id)
);

CREATE TABLE order_detail (
    order_detail_id SERIAL PRIMARY KEY,
    order_id INT NOT NULL ,
    service_id INT NOT NULL,
    qty INT NOT NULL,
    FOREIGN KEY(order_id) REFERENCES "order"(order_id),
    FOREIGN KEY(service_id) REFERENCES service(service_id)
);


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
