package main

import (
	"database/sql"
	"fmt"

	"project-golangDB/entity"

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

func main() {
	customer := entity.Customer{Id: 3, Name: "king", Phone: "0819774393", Address: "Bandung"}
	create_Customer(customer)
	// customers := view_Of_List_Customer()
	// for _, customer := range customers {
	// 	fmt.Println(customer.Id, customer.Name, customer.Phone, customer.Address, customer.Created_at, customer.Updated_at)
	// }
	// fmt.Println(get_Customer_By_ID(5))
	// update_Customer(customer)
	// delete_Customer(6)

}

// func enroll_Customer(customer_Enrollment entity.Customer){
// 	db := connect_Db()
// 	defer db.Close()

// 	tx, err := db.Begin()
// 	if err != nil {
// 		panic(err)
// 	}

// 	create_Customer(customer_Enrollment, tx)

// 	update_Order(customer_Enrollment.Id, tx)

// 	if err := tx.Commit(); err != nil {
// 		fmt.Println("Failed to commit transaction:", err)
// 		tx.Rollback() // rollback jika commit gagal
// 	} else {
// 		fmt.Println("Transaction committed successfully")
// 	}
// }

// func update_Order(customer_Id int, tx *sql.Tx) {
// 	update_Order := "UPDATE \"order\" SET customer_id = $1"

// 	_,err := tx.Exec(update_Order, customer_Id)
// 	validate(err, "Update", tx)
// }

func create_Customer(customer entity.Customer) {
	db := connect_Db()
	defer db.Close()
	var err error

	if check_Query(db, customer.Id) {
		fmt.Println("Customer ID already exists. Please enter a different ID.")
		return
	}

	insert_Customer := "INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4);"

	_ , err = db.Exec(insert_Customer, customer.Id, customer.Name, customer.Phone, customer.Address)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
	// validate(err, "Insert", tx)
}

func view_Of_List_Customer() []entity.Customer {
	db := connect_Db()
	defer db.Close()

	sql_Statement := "SELECT * FROM customer"

	rows, err := db.Query(sql_Statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := scan_Customers(rows)

	return customers
}

func scan_Customers(rows *sql.Rows) []entity.Customer {
	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return customers
}

func get_Customer_By_ID(id int) entity.Customer {
	db := connect_Db()
	defer db.Close()

	var err error

	sql_Statement :=  "SELECT * FROM customer WHERE customer_id = $1"
	
	customer := entity.Customer{}
	if !check_Query(db, id) {
		fmt.Println("Customer not found.")
	} 

	err = db.QueryRow(sql_Statement, id).Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address, &customer.Created_at, &customer.Updated_at)
	if err != nil {
		panic(err)
	}
	
	return customer
}

func update_Customer(customer entity.Customer) {
	db := connect_Db()
	defer db.Close()
	var err error

	if !check_Query(db, customer.Id) {
		fmt.Println("Customer not found")
		return
	}

	sql_Statement := "UPDATE customer SET name = $2, phone = $3, address = $4,  updated_at = CURRENT_TIMESTAMP WHERE customer_id = $1;"

	_, err = db.Exec(sql_Statement, customer.Id, customer.Name, customer.Phone, customer.Address)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func delete_Customer(id int) {
	db := connect_Db()
	defer db.Close()
	var err error 

	if !check_Query(db, id) {
		fmt.Println("Customer not found")
		return
	}

	sql_Statement := "DELETE FROM customer WHERE customer_id = $1;"

	_, err = db.Exec(sql_Statement, id)
	if err != nil {
		panic(err)		
	} else {
		fmt.Println("Successfully Delete Data!")
	}
}

func create_Services(service entity.Service) {
	db := connect_Db()
	defer db.Close()
	var err error
	if check_Query(db, service.Id) {
		fmt.Println("Service ID already exists. Please enter a different ID.")
		return
	}

	insert_Service := "INSERT INTO service(service_id, service_name, unit, price) VALUES ($1, $2, $3, $4);"

	_, err = db.Exec(insert_Service, service.Id, service.Name, service.Unit, service.Price)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func  view_Of_List_Service() []entity.Service {
	db := connect_Db()
	defer db.Close()

	view_Service := "SELECT * FROM service;"

	rows, err := db.Query(view_Service)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	view_Services := scan_Service(rows)

	return view_Services
}

func scan_Service(rows *sql.Rows) []entity.Service {
	services := []entity.Customer{}
	var err error

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Id, &service.Name, &service.Unit, &service.Price, &service.Created_at, service.Updated_at)
		if err != nil{
			panic(err)
		}

		services = append(services, service)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}

func view_Details_Service_By_ID(id int) entity.Service{
	db := connect_Db()
	defer db.Close()
	var err error

	view_By_ID := "SELECT * FROM service WHERE id = $1;"

	service := entity.Service{}
	err = db.QueryRow(view_By_ID, id).Scan(&service.Id, &service.Name, &service.Price, &service.Unit, service.Created_at, service.Updated_at)
	if err != nil {
		panic(err)
	}

	return service
}

func update_Service(service entity.Service){
	db := connect_Db()
	defer db.Close()
	var err error

	sql := "UPDATE service SET service_name = $2, unit = $3, price = $4, updated_at = $5 WHERE service_id = $1;"

	_, err = db.Exec(sql, service.Id, service.Name, service.Unit, service.Price, service.Updated_at)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Update Data!")
	}
}

func Delete_Service(id int) {
	db := connect_Db()
	defer db.Close()
	var err error

	sql := "DELETE FROM service WHERE service_id = $1;"

	_, err = db.Exec(sql, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Data!.")
	}
}

func create_Order(order entity.Order) {
	db := connect_Db()
	defer db.Close()
	var err error
	if check_Query(db, order.Id) {
		fmt.Println("Customer ID already exists. Please enter a different ID.")
		return
	}

	insert_Order := "INSERT INTO  \"order\" (order_id, customer_id,  received_by) VALUES ($1, $2, $3);"

	_, err = db.Exec(insert_Order, order.Id, order.Customer_id, order.Order_date, order.Received_by)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data Order!")
	}
}

func Complete_Order(order entity.Order){
	db := connect_Db()
	defer db.Close()
	var err error

	sql_Statement := "UPDATE  \"order\" SET completion_date = $2,  updated_date = $3 WHERE id = $1;"

	_, err = db.Exec(sql_Statement, order.Id, order.Completion_date, order.Updated_at )
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Complete Order!")
	}
}

func View_Of_List_Order() []entity.Order {
	db := connect_Db()
	defer db.Close()

	sql_Statement := "SELECT * FROM \"order\";"

	rows, err := db.Query(sql_Statement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orders := scan_Order(rows)

	return orders
}

func scan_Order(rows *sql.Rows) []entity.Order {
	orders := []entity.Order{}
	var err error

	for rows.Next() {
		order := entity.Order{}
		err := rows.Scan(&order.Id, &order.Customer_id,&order.Order_date, &order.Completion_date, order.Received_by, order.Created_at, &order.Updated_at)
		if err != nil{
			panic(err)
		}

		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return orders
}

func View_Order_Details_By_ID(id int) entity.Order{
	db := connect_Db()
	defer db.Close()
	var err error

	sql_Statement := "SELECT * FORM \"order\" WHERE order_id = $1;"

	order := entity.Order{}
	err = db.QueryRow(sql_Statement, id).Scan(&order.Id, &order.Customer_id, &order.Order_date, &order.Completion_date, &order.Received_by, &order.Created_at, &order.Updated_at)
	if err != nil {
		panic(err)
	}

	return order
}

func check_Query(db *sql.DB, customer_Id int) bool {
	var exists bool
	checkQuery := "SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)"
	err := db.QueryRow(checkQuery, customer_Id).Scan(&exists)
	if err != nil {
		panic(err)
	}

	return exists
}

// func check_Query1(tx *sql.Tx, customer_Id int) bool {
// 	var exists bool
// 	checkQuery := "SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)"
// 	err := tx.QueryRow(checkQuery, customer_Id).Scan(&exists)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return exists
// }

// func validate(err error, messsage string, tx *sql.Tx){
// 	if err != nil {
// 		tx.Rollback()
// 		fmt.Println(err, "Transaction Rollback!")
// 	} else {
// 		fmt.Println("Succesfully " + messsage + " data!")
// 	}
// }

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
