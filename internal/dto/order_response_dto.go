package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"
)

type UpdateOrdersByIDResponse struct {
	CustomerName string    `json:"customer_name"`
	UpdatedAt    time.Time `json:"updated_at"`
	ItemsUpdate  `json:"items"`
}

type ItemUpdateResponse struct {
	ItemID      uint64 `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
}

type PersonResponse struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Username      string `json:"username"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Uuid          string `json:"uuid"`
	OrderResponse `json:"orders"`
}

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

type ViewOrderResponse struct {
	OrderID            uint64    `json:"order_id"`
	CustomerName       string    `json:"customer_name"`
	CreatedAt          time.Time `json:"created_at"`
	ViewItemsResponses `json:"items"`
}

type ViewItemsResponse struct {
	ItemID      uint64 `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
}

type ViewOrderResponses []ViewOrderResponse
type ViewItemsResponses []ViewItemsResponse
type OrderDetails []OrderResponse
type AllItems []ItemsResponse
type ItemsUpdate []ItemUpdateResponse

func CreateItemUpdateResponse(i entity.Items) ItemUpdateResponse {
	return ItemUpdateResponse{
		ItemID:      i.ItemId,
		ItemCode:    i.ItemCode,
		Description: i.Description,
		Quantity:    i.Quantity,
	}
}

func CreatePersonOrderResponse(p entity.Person) PersonResponse {
	return PersonResponse{
		FirstName: p.Result[0].Firstname,
		LastName:  p.Result[0].Lastname,
		Username:  p.Result[0].Username,
		Email:     p.Result[0].Email,
		Phone:     p.Result[0].Phone,
		Uuid:      p.Result[0].UUID,
	}
}
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

func CreateViewItemsResponse(item entity.Items) ViewItemsResponse {
	return ViewItemsResponse{
		ItemID:      item.ItemId,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
	}
}

func CreateUpdateResponse(o entity.Orders, is entity.AllItems) *UpdateOrdersByIDResponse {
	updateDetails := UpdateOrdersByIDResponse{}
	updateDetails.CustomerName = o.CustomerName
	updateDetails.UpdatedAt = time.Now()
	for idx := range is {
		allItems := CreateItemUpdateResponse(*is[idx])
		updateDetails.ItemsUpdate = append(updateDetails.ItemsUpdate, allItems)
	}
	return &updateDetails
}

func CreatePersonOrdersResponse(os entity.OrdersItemsJoined, p entity.Person) *PersonResponse {
	personOrdersDetails := CreatePersonOrderResponse(p)
	personOrdersDetails.OrderResponse = viewPersonOrderResponse(os)
	return &personOrdersDetails
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

func ViewOrderResponseDetails(os entity.OrdersJoined) []ViewOrderResponse {
	var orderDetails []ViewOrderResponse

	for _, each := range os {
		order := viewOrderResponse(*each)
		orderDetails = append(orderDetails, order)
	}
	return orderDetails
}

func viewOrderResponse(os entity.OrdersItemsJoined) ViewOrderResponse {
	var listItems ViewItemsResponses

	for _, each := range os.Items {
		items := CreateViewItemsResponse(*each)
		listItems = append(listItems, items)
	}

	return ViewOrderResponse{
		OrderID:            os.OrderID,
		CustomerName:       os.CustomerName,
		CreatedAt:          os.CreatedAt,
		ViewItemsResponses: listItems,
	}
}

func viewPersonOrderResponse(os entity.OrdersItemsJoined) OrderResponse {
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
