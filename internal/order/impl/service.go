package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
)

type orderServiceImpl struct {
	repo OrderRepository
}

func ProvideOrderService(repo OrderRepository) *orderServiceImpl {
	return &orderServiceImpl{
		repo: repo,
	}
}

func (o orderServiceImpl) CreateNewOrder(ctx context.Context, data *dto.CreateOrderRequest) (*dto.OrderResponse, error) {
	item, order := data.ToEntity()
	orderID, err := o.repo.CreateNewOrder(ctx, order, item)
	if err != nil {
		log.Printf("[CreateNewOrder] an error occured while creating new order, err => %v", err)
		return nil, err
	}

	return dto.CreateOrderResponseDetail(order, item, orderID), nil
}

func (o orderServiceImpl) ViewAllOrders(ctx context.Context) (*dto.OrderDetails, error) {
	res, err := o.repo.ViewAllOrders(ctx)
	if err != nil {
		log.Printf("[ViewAllOrders] an error occured while show all the orders, err => %v", err)
		return nil, err
	}

	var response1 dto.OrderDetails
	response := dto.ViewOrderResponseDetails(res)
	response1 = append(response1, response...)

	return &response1, nil
}
