package order

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/order/impl"
)

type OrderService interface {
	CreateNewOrder(ctx context.Context, data *dto.CreateOrderRequest) (*dto.OrderResponse, error)
	ViewAllOrders(ctx context.Context) (*dto.OrderDetails, error)
	DeleteOrderByID(ctx context.Context, id uint64) (int, error)
	UpdateOrderByID(ctx context.Context, id uint64, data *dto.UpdateOrderRequest) (*dto.UpdateOrdersByIDResponse, error)
	PersonOrders(ctx context.Context, id uint64) (*dto.PersonResponse, error)
}

func ProvideOrderService(db *sql.DB) OrderService {
	repo := impl.ProvideOrderRepository(db)
	return impl.ProvideOrderService(repo)
}
