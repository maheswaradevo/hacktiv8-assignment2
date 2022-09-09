package impl

import "database/sql"

type orderRepositoryImpl struct {
	db *sql.DB
}

func ProvideOrderRepository(db *sql.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{
		db: db,
	}
}

type OrderRepository interface {
}
