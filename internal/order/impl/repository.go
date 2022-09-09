package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"
)

type orderRepositoryImpl struct {
	db *sql.DB
}

func ProvideOrderRepository(db *sql.DB) *orderRepositoryImpl {
	return &orderRepositoryImpl{
		db: db,
	}
}

type OrderRepository interface {
	CreateNewOrder(ctx context.Context, reqDataOrder entity.Orders, reqDataItems entity.AllItems) (uint64, error)
}

var (
	INSERT_ORDER_DATA = "INSERT INTO `order` (customer_name) VALUES (?)"
	INSERT_ITEM_DATA  = "INSERT INTO `item` (item_code, description, quantity, order_id) VALUES(?, ?, ?, ?)"
)

func (o orderRepositoryImpl) CreateNewOrder(ctx context.Context, reqDataOrder entity.Orders, reqDataItems entity.AllItems) (uint64, error) {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[CreateNewOrder] failed to begin transaction, err => %v", err)
		return 0, err
	}
	defer tx.Rollback()

	queryOrder := INSERT_ORDER_DATA
	res, err := tx.ExecContext(ctx, queryOrder, reqDataOrder.CustomerName)
	if err != nil {
		log.Printf("[CreateNewOrder] failed to insert order data, err => %v", err)
		return 0, err
	}

	orderID, _ := res.LastInsertId()

	for _, items := range reqDataItems {
		queryItemData := INSERT_ITEM_DATA
		stmt, err := tx.PrepareContext(ctx, queryItemData)
		if err != nil {
			log.Printf("[CreateNewOrder] failed to prepare the statement, err => %v", err)
			return 0, err
		}
		res, err = stmt.ExecContext(
			ctx,
			items.ItemCode,
			items.Description,
			items.Quantity,
			orderID,
		)
		if err != nil {
			log.Printf("[CreateNewOrder] failed to insert items data, err => %v", err)
			return 0, err
		}
		itemsID, _ := res.LastInsertId()
		items.ItemId = uint64(itemsID)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[CreateNewOrder] transaction failed, err => %v", err)
		return 0, err
	}
	return uint64(orderID), nil
}
