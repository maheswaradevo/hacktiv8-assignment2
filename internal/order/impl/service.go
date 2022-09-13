package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/utils"
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
	count, err := o.repo.CheckOrders(ctx)
	if err != nil {
		log.Printf("[ViewAllOrders] an error occured while show all the orders, err => %v", err)
		return nil, err
	}
	if count < 1 {
		log.Printf("[ViewAllOrders] theres is no orders data, err => %v", err)
		panic(
			utils.NewErrorResponse(
				494,
				"DATA_NOT_EXISTS",
				utils.NewErrorResponseValue("orders", "does not exists"),
			),
		)

	}
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
