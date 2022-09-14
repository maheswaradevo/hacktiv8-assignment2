package dto

import (
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"
)

type UpdateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	ItemsRequest `json:"items"`
}
type CreateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	ItemsRequest `json:"items"`
}

type ItemRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
}

type ItemsRequest []ItemRequest

func (dto *CreateOrderRequest) ToEntity() (item entity.AllItems, order entity.Orders) {
	item = entity.AllItems{}
	for _, items := range dto.ItemsRequest {
		itemsDetail := entity.Items{
			ItemCode:    items.ItemCode,
			Description: items.Description,
			Quantity:    items.Quantity,
		}
		item = append(item, &itemsDetail)
	}

	order = entity.Orders{
		CustomerName: dto.CustomerName,
	}
	return
}

func (dto *UpdateOrderRequest) ToEntity() (order entity.Orders, item entity.AllItems) {
	order = entity.Orders{
		CustomerName: dto.CustomerName,
	}

	item = entity.AllItems{}
	for _, items := range dto.ItemsRequest {
		itemsDetails := entity.Items{
			ItemCode:    items.ItemCode,
			Description: items.Description,
			Quantity:    items.Quantity,
		}
		item = append(item, &itemsDetails)
	}

	return
}
