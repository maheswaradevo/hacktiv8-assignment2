package order

import (
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/order/impl"
)

type OrderService interface {
}

func ProvideOrderService(db *sql.DB) OrderService {
	repo := impl.ProvideOrderRepository(db)
	return impl.ProvideOrderService(repo)
}
