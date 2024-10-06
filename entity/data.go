package entity

import "time"

type Customer struct {
	Id        int
	Name      string
	Phone     string
	Address   string
	Created_at time.Time
	Updated_at time.Time
}

type Service struct {
	Id        int
	Name      string
	Unit string
	Price     int
	Created_at time.Time
	Updated_at time.Time
}

type Order struct {
	Id        int
	Customer_id int
	Order_date time.Time
	Completion_date time.Time
	Received_by string
	Created_at time.Time
	Updated_at time.Time
}

type Order_detail struct {
	Order_detail_id int
	Order_id int
	Service_id int
	Qty int
}