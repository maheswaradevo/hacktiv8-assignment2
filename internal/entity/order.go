package entity

import "time"

type Orders struct {
	OrderID      uint64    `db:"order_id"`
	CustomerName string    `db:"customer_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type Items struct {
	ItemId      uint64    `db:"item_id"`
	ItemCode    string    `db:"item_code"`
	Description string    `db:"description"`
	Quantity    uint64    `db:"quantity"`
	OrderID     uint64    `db:"order_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type OrdersItemsJoined struct {
	Orders
	Items
}

type OrdersJoined []*OrdersItemsJoined
type AllItems []*Items
type AllOrders []*Orders
