package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"
)

type OrderResponse struct {
	OrderID      uint64    `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	CreatedAt    time.Time `json:"ordered_at"`
	AllItems     `json:"items"`
}

type ItemsResponse struct {
	ItemID      uint64 `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	OrderID     uint64 `json:"order_id"`
}

type OrderDetails []OrderResponse
type AllItems []ItemsResponse

func CreateOrderResponse(order entity.Orders) OrderResponse {
	return OrderResponse{
		CustomerName: order.CustomerName,
		CreatedAt:    time.Now(),
	}
}

func CreateItemsResponse(item entity.Items) ItemsResponse {
	return ItemsResponse{
		ItemID:      item.ItemId,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}

func CreateOrderResponseDetail(o entity.Orders, is entity.AllItems, id uint64) *OrderResponse {
	orderDetails := CreateOrderResponse(o)
	orderDetails.OrderID = id
	for idx := range is {
		items := CreateItemsResponse(*is[idx])
		items.OrderID = id
		orderDetails.AllItems = append(orderDetails.AllItems, items)
	}

	return &orderDetails
}

func viewOrderResponse(os entity.OrdersItemsJoined) OrderResponse {
	var listItems AllItems

	for _, each := range os.Items {
		items := CreateItemsResponse(*each)
		listItems = append(listItems, items)
	}

	return OrderResponse{
		OrderID:      os.OrderID,
		CustomerName: os.CustomerName,
		CreatedAt:    os.CreatedAt,
		AllItems:     listItems,
	}
}

func ViewOrderResponseDetails(os entity.OrdersJoined) []OrderResponse {
	var orderDetails []OrderResponse

	for _, each := range os {
		order := viewOrderResponse(*each)
		orderDetails = append(orderDetails, order)
	}
	return orderDetails
}
