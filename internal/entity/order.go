package entity

import "time"

type Orders struct {
	OrderID      uint64
	CustomerName string
	OrderedAt    time.Time
}

type Items struct {
	ItemId      uint64
	ItemCode    string
	Description string
	Quantity    uint64
	OrderID     uint64
}
