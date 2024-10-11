package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "enigma_laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

var db *sql.DB

func main() {
	db = connect_Db()  // Connect to the database
	defer db.Close()    // Make sure to close the connection when the program ends
	mainMenu()
}

func mainMenu() {
    var choice int
	for {
    fmt.Println("Main Menu:")
    fmt.Println("1. Customer")
    fmt.Println("2. Service")
    fmt.Println("3. Order")
    fmt.Println("4. Exit")
    fmt.Print("Enter choice: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        customerMenu()
    case 2:
        serviceMenu()
    case 3:
        orderMenu()
    case 4:
        fmt.Println("Exiting...")
        return
    default:
        fmt.Println("Invalid choice.")
        mainMenu()
    }
}
}

func customerMenu() {
    var choice int
    fmt.Println("Customer Menu:")
    fmt.Println("1. Create Customer")
    fmt.Println("2. View Customer List")
    fmt.Println("3. View Customer Details")
    fmt.Println("4. Update Customer")
    fmt.Println("5. Delete Customer")
    fmt.Println("6. Back to Main Menu")
    fmt.Print("Enter choice: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        createCustomer()
    case 2:
        viewCustomerList()
    case 3:
        viewCustomerDetails()
    case 4:
        updateCustomer()
    case 5:
        deleteCustomer()
    case 6:
        mainMenu()
    default:
        fmt.Println("Invalid choice.")
        customerMenu()
    }
}

func createCustomer() {
    var id int
    var name, phone, address string
    fmt.Print("Enter customer ID: ")
    fmt.Scanln(&id)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customer WHERE customer_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if exists {
        fmt.Println("Customer ID already exists. Please enter a different ID.")
        return
    }

    fmt.Print("Enter name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter phone: ")
    fmt.Scanln(&phone)
    fmt.Print("Enter address: ")
    fmt.Scanln(&address)

    _, err = db.Exec("INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4)", id, name, phone, address)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Customer created successfully.")
}

func viewCustomerList() {
    rows, err := db.Query("SELECT customer_id, name, phone, address FROM customer")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Customer List:")
    for rows.Next() {
        var id int
        var name, phone, address string
        err := rows.Scan(&id, &name, &phone, &address)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Phone: %s, Address: %s\n", id, name, phone, address)
    }
}

func viewCustomerDetails() {
    var id int
    fmt.Print("Enter customer ID: ")
    fmt.Scanln(&id)

    var name, phone, address string
    err := db.QueryRow("SELECT name, phone, address FROM customer WHERE customer_id=$1", id).Scan(&name, &phone, &address)
    if err == sql.ErrNoRows {
		fmt.Println("Customer not found.")
		return
	} else if err != nil {
		log.Fatal(err)
	}	

    fmt.Printf("ID: %d, Name: %s, Phone: %s, Address: %s\n", id, name, phone, address)
}

func updateCustomer() {
    var id int
    var name, phone, address string
    fmt.Print("Enter customer ID: ")
    fmt.Scanln(&id)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customer WHERE customer_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if !exists {
        fmt.Println("Customer not found.")
        return
    }

    fmt.Print("Enter new name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter new phone: ")
    fmt.Scanln(&phone)
    fmt.Print("Enter new address: ")
    fmt.Scanln(&address)

    _, err = db.Exec("UPDATE customer SET name=$1, phone=$2, address=$3, updated_at=NOW() WHERE customer_id=$4", name, phone, address, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Customer updated successfully.")
}

func deleteCustomer() {
    var id int
    fmt.Print("Enter customer ID: ")
    fmt.Scanln(&id)

    // Cek apakah customer ID ada di database
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM customer WHERE customer_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }

    if !exists {
        // Jika ID tidak ditemukan
        fmt.Println("Customer ID not found. Please enter a different ID.")
        return
    }

    // Cek apakah customer ID sudah digunakan dalam tabel order
    var order_Exists bool
    err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM \"order\" WHERE customer_id=$1)", id).Scan(&order_Exists)
    if err != nil {
        log.Fatal(err)
    }

    if order_Exists {
        // Jika customer ID digunakan di dalam order
        fmt.Println("Customer ID is being used in orders. Please delete the order first.")
        return
    }

    // Jika customer tidak terlibat di order, hapus customer
    _, err = db.Exec("DELETE FROM customer WHERE customer_id=$1", id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Customer deleted successfully.")
}

func serviceMenu() {
    var choice int
    fmt.Println("Service Menu:")
    fmt.Println("1. Create Service")
    fmt.Println("2. View Service List")
    fmt.Println("3. View Service Details")
    fmt.Println("4. Update Service")
    fmt.Println("5. Delete Service")
    fmt.Println("6. Back to Main Menu")
    fmt.Print("Enter choice: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        createService()
    case 2:
        viewServiceList()
    case 3:
        viewServiceDetails()
    case 4:
        updateService()
    case 5:
        deleteService()
    case 6:
        mainMenu()
    default:
        fmt.Println("Invalid choice.")
        serviceMenu()
    }
}

func createService() {
    var id int
    var name, unit string
    var price int

    fmt.Print("Enter service ID: ")
    fmt.Scanln(&id)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM service WHERE service_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if exists {
        fmt.Println("Service ID already exists. Please enter a different ID.")
        return
    }

    fmt.Print("Enter service name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter unit: ")
    fmt.Scanln(&unit)
    fmt.Print("Enter price: ")
    fmt.Scanln(&price)

    _, err = db.Exec("INSERT INTO service (service_id, service_name, unit, price) VALUES ($1, $2, $3, $4)", id, name, unit, price)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Service created successfully.")
}

func viewServiceList() {
    rows, err := db.Query("SELECT service_id, service_name, unit, price FROM service")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Service List:")
    for rows.Next() {
        var id, price int
        var name, unit string
        err := rows.Scan(&id, &name, &unit, &price)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Name: %s, Unit: %s, Price: %d\n", id, name, unit, price)
    }
}

func viewServiceDetails() {
    var id int
    fmt.Print("Enter service ID: ")
    fmt.Scanln(&id)

    var name, unit string
    var price int
    err := db.QueryRow("SELECT service_name, unit, price FROM service WHERE service_id=$1", id).Scan(&name, &unit, &price)
    if err == sql.ErrNoRows {
        fmt.Println("Service not found.")
        return
    } else if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("ID: %d, Name: %s, Unit: %s, Price: %d\n", id, name, unit, price)
}

func updateService() {
    var id int
    var name, unit string
    var price int

    fmt.Print("Enter service ID: ")
    fmt.Scanln(&id)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM service WHERE service_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if !exists {
        fmt.Println("Service not found.")
        return
    }

    fmt.Print("Enter new service name: ")
    fmt.Scanln(&name)
    fmt.Print("Enter new unit: ")
    fmt.Scanln(&unit)
    fmt.Print("Enter new price: ")
    fmt.Scanln(&price)

    _, err = db.Exec("UPDATE service SET service_name=$1, unit=$2, price=$3, updated_at=NOW() WHERE service_id=$4", name, unit, price, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Service updated successfully.")
}

func deleteService() {
    var id int
    fmt.Print("Enter service ID: ")
    fmt.Scanln(&id)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM service WHERE service_id=$1)", id).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if !exists {
        fmt.Println("Service ID not found. Please enter a different ID.")
        return
    }

    // Check if the service is used in any orders
    var order_Exists bool
    err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM order_detail WHERE service_id=$1)", id).Scan(&order_Exists)
    if err != nil {
        log.Fatal(err)
    }
    if order_Exists {
        fmt.Println("Service ID is being used in orders. Please delete the order first.")
        return
    }

    _, err = db.Exec("DELETE FROM service WHERE service_id=$1", id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Service deleted successfully.")
}

func orderMenu() {
    var choice int
    fmt.Println("Order Menu:")
    fmt.Println("1. Create Order")
    fmt.Println("2. Complete Order")
    fmt.Println("3. View Order List")
    fmt.Println("4. View Order Details")
    fmt.Println("5. Back to Main Menu")
    fmt.Print("Enter choice: ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        createOrder()
    case 2:
        completeOrder()
    case 3:
        viewOrderList()
    case 4:
        viewOrderDetails()
    case 5:
        mainMenu()
    default:
        fmt.Println("Invalid choice.")
        orderMenu()
    }
}

func createOrder() {
    var orderID, customerID int
    var receivedBy string
    var serviceID, quantity int

    // Input order information
    fmt.Print("Enter order ID: ")
    fmt.Scanln(&orderID)

    // Cek apakah order ID sudah ada
    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM \"order\" WHERE order_id=$1)", orderID).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if exists {
        fmt.Println("Order ID already exists. Please enter a different ID.")
        return
    }

    // Input customer ID
    fmt.Print("Enter customer ID: ")
    fmt.Scanln(&customerID)

    err = db.QueryRow("SELECT EXISTS (SELECT 1 FROM customer WHERE customer_id=$1)", customerID).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if !exists {
        fmt.Println("Customer not found.")
        return
    }

    // Input service details
    fmt.Print("Enter received by: ")
    fmt.Scanln(&receivedBy)
    fmt.Print("Enter service ID: ")
    fmt.Scanln(&serviceID)
    fmt.Print("Enter quantity: ")
    fmt.Scanln(&quantity)

    // Mulai transaksi
    tx, err := db.Begin()
    if err != nil {
        log.Fatal("Failed to begin transaction:", err)
    }

    // Insert ke tabel order
    _, err = tx.Exec("INSERT INTO \"order\" (order_id, customer_id, order_date, received_by) VALUES ($1, $2, NOW(), $3)", orderID, customerID, receivedBy)
    if err != nil {
        // Jika gagal, rollback transaksi
        tx.Rollback()
        fmt.Println("Failed to insert into order:", err)
        return
    }

    // Insert ke tabel order_detail
    _, err = tx.Exec("INSERT INTO order_detail (order_id, service_id, qty)VALUES ($1, $2, $3)", orderID, serviceID, quantity)
    if err != nil {
        // Jika gagal, rollback transaksi
        tx.Rollback()
        fmt.Println("Failed to insert into order_detail:", err)
        return
    }

    // Commit transaksi jika semua berhasil
    err = tx.Commit()
    if err != nil {
        fmt.Println("Failed to commit transaction:", err)
    } else {
        fmt.Println("Order and order detail created successfully.")
    }
}

func completeOrder() {
    var orderID int
    var completionDate string

    fmt.Print("Enter order ID: ")
    fmt.Scanln(&orderID)

    var exists bool
    err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM \"order\" WHERE order_id=$1)", orderID).Scan(&exists)
    if err != nil {
        log.Fatal(err)
    }
    if !exists {
        fmt.Println("Order not found.")
        return
    }

    fmt.Print("Enter completion date (YYYY-MM-DD): ")
    fmt.Scanln(&completionDate)

    _, err = db.Exec("UPDATE \"order\" SET completion_date=$1, updated_at=NOW() WHERE order_id=$2", completionDate, orderID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Order completed successfully.")
}

func viewOrderList() {
    rows, err := db.Query("SELECT order_id, customer_id, order_date, received_by FROM \"order\"")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Println("Order List:")
    for rows.Next() {
        var order_ID, customer_ID int
        var order_Date, received_By string
        err := rows.Scan(&order_ID, &customer_ID, &order_Date, &received_By)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("Order ID: %d, Customer ID: %d, Order Date: %s, Received By: %s\n", order_ID, customer_ID, order_Date, received_By)
    }
}

func viewOrderDetails() {
    var order_ID int
    fmt.Print("Enter order ID: ")
    fmt.Scanln(&order_ID)

    var customer_ID int
    var order_Date, received_By, completion_Date string
    err := db.QueryRow("SELECT customer_id, order_date, received_by, completion_date FROM \"order\" WHERE order_id=$1", order_ID).Scan(&customer_ID, &order_Date, &received_By, &completion_Date)
    if err == sql.ErrNoRows {
        fmt.Println("Order not found.")
        return
    } else if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Order ID: %d, Customer ID: %d, Order Date: %s, Received By: %s, Completion Date: %s\n", order_ID, customer_ID, order_Date, received_By, completion_Date)
}

func connect_Db() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Connected!")
	}

	return db
}