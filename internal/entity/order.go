package entity

import "time"

// Orders represent the model for an order
type Orders struct {
	OrderID      uint64    `db:"order_id"`
	CustomerName string    `db:"customer_name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

// Items represent the model for an item in the order
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
	Items AllItems
}

type OrdersJoined []*OrdersItemsJoined
type AllItems []*Items
type AllOrders []*Orders
