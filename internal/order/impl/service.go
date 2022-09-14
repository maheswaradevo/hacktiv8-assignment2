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

func (o orderServiceImpl) UpdateOrderByID(ctx context.Context, id uint64, data *dto.UpdateOrderRequest) (int, error) {
	check, err := o.repo.CheckOrders(ctx)
	if err != nil {
		log.Printf("[UpdateOrderByID] an error occured while checking orders, err => %v, id => %v", err, id)
		return 0, nil
	}
	if check < 1 {
		log.Printf("[DeleteOrderByID] theres is no orders data, err => %v", err)
		panic(err)
	}

	order, item := data.ToEntity()

	res, err := o.repo.UpdateOrderByID(ctx, id, &order, item)
	if err != nil {
		log.Printf("[UpdateOrderByID] an error occured while updating orders, err => %v, id => %v", err, id)
		return 0, nil
	}

	return res, nil
}

func (o orderServiceImpl) DeleteOrderByID(ctx context.Context, id uint64) (int, error) {
	check, err := o.repo.CheckOrders(ctx)
	if err != nil {
		log.Printf("[DeleteOrderByID] an error occured while checking orders, err => %v, id => %v", err, id)
		return 0, nil
	}
	if check < 1 {
		log.Printf("[DeleteOrderByID] theres is no orders data, err => %v", err)
		panic(err)
	}
	res, err := o.repo.DeleteOrderByID(ctx, id)
	if err != nil {
		log.Printf("[DeleteOrderByID] an error occured while deleting orders, err => %v, id => %v", err, id)
		return 0, nil
	}

	return res, nil
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
		panic(err)

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
