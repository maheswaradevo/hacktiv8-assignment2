package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

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
	ViewAllOrders(ctx context.Context) (entity.OrdersJoined, error)
	CheckOrders(ctx context.Context) (int, error)
	DeleteOrderByID(ctx context.Context, id uint64) (int, error)
	UpdateOrderByID(ctx context.Context, id uint64, reqDataOrder *entity.Orders, reqDataItems entity.AllItems) error
	FetchPerson() (entity.Person, error)
	GetOrdersByID(ctx context.Context, id uint64) (entity.OrdersItemsJoined, error)
}

var (
	INSERT_ORDER_DATA   = "INSERT INTO `order` (customer_name) VALUES (?)"
	INSERT_ITEM_DATA    = "INSERT INTO `item` (item_code, description, quantity, order_id) VALUES(?, ?, ?, ?)"
	SELECT_ORDERS       = "SELECT o.order_id, o.customer_name, o.created_at, o.updated_at FROM `order` o"
	SELECT_ORDERS_BY_ID = "SELECT o.order_id, o.customer_name, o.created_at, o.updated_at FROM `order` o WHERE o.order_id = ?"
	SELECT_ITEMS        = "SELECT i.item_id, i.item_code, i.description, i.quantity, i.order_id FROM `item` i WHERE i.order_id=?"
	COUNT_ORDERS        = "SELECT COUNT(*) FROM `order`"
	COUNT_ITEMS         = "SELECT COUNT(*) FROM `item` i WHERE i.order_id = ?"
	DELETE_ORDER        = "DELETE FROM `order` WHERE order_id = ?"
	DELETE_ITEMS        = "DELETE FROM `item` WHERE order_id = ?"
	UPDATE_ORDER        = "UPDATE `order` SET customer_name = ? WHERE order_id = ?"
	UPDATE_ITEM         = "UPDATE item SET description = ?, item_code = ?, quantity = ? WHERE item_id = ?"
)

var baseUrl = "https://hiyaa.site/data.php?qty=2&apikey=7f8fc96e-de1f-4aab-9c62-3dd1de365e66"

func (o orderRepositoryImpl) GetOrdersByID(ctx context.Context, id uint64) (entity.OrdersItemsJoined, error) {
	query := SELECT_ORDERS_BY_ID

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[GetOrdersByID] failed to begin transaction, err => %v", err)
		return entity.OrdersItemsJoined{}, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetOrdersByID] failed to prepare the statement, err => %v", err)
		return entity.OrdersItemsJoined{}, err
	}

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		log.Printf("[GetOrdersByID] failed to query to the database, err => %v", err)
		return entity.OrdersItemsJoined{}, err
	}

	var orders entity.OrdersItemsJoined
	for rows.Next() {
		orderDetails := entity.OrdersItemsJoined{}

		err := rows.Scan(
			&orderDetails.Orders.OrderID,
			&orderDetails.Orders.CustomerName,
			&orderDetails.Orders.CreatedAt,
			&orderDetails.Orders.UpdatedAt,
		)
		if err != nil {
			log.Printf("[GetOrdersByID] failed to scan the data, err => %v", err)
			return entity.OrdersItemsJoined{}, err
		}
		orders = orderDetails
	}
	countItems := COUNT_ITEMS

	stmt, err = o.db.PrepareContext(ctx, countItems)
	if err != nil {
		log.Printf("[GetOrdersByID] failed to prepare the statement, err => %v", err)
		return entity.OrdersItemsJoined{}, err
	}

	rows, err = stmt.QueryContext(ctx, id)
	if err != nil {
		log.Printf("[GetOrdersByID] failed to query to the database, err => %v", err)
		return entity.OrdersItemsJoined{}, err
	}

	var itemCount int

	for rows.Next() {
		err := rows.Scan(
			&itemCount,
		)
		if err != nil {
			log.Printf("[GetOrdersByID] failed to scan the data, err => %v", err)
			return entity.OrdersItemsJoined{}, err
		}
	}
	queryItems := SELECT_ITEMS
	for idx := 1; idx < itemCount; idx++ {
		stmt, err := tx.PrepareContext(ctx, queryItems)
		if err != nil {
			log.Printf("[GetOrdersByID] failed to prepare the statement, err => %v", err)
			return entity.OrdersItemsJoined{}, err
		}

		rows, err := stmt.QueryContext(ctx, orders.OrderID)
		if err != nil {
			log.Printf("[GetOrdersByID] failed to query to the database, err => %v", err)
			return entity.OrdersItemsJoined{}, err
		}

		for rows.Next() {
			allItems := entity.Items{}

			err := rows.Scan(
				&allItems.ItemId,
				&allItems.ItemCode,
				&allItems.Description,
				&allItems.Quantity,
				&allItems.OrderID,
			)
			if err != nil {
				log.Printf("[GetOrdersByID] failed to scan the data, err => %v", err)
				return entity.OrdersItemsJoined{}, err
			}
			orders.Items = append(orders.Items, &allItems)
		}
	}
	return orders, nil
}

func (o orderRepositoryImpl) FetchPerson() (entity.Person, error) {
	client := &http.Client{}

	people := entity.Person{}

	request, err := http.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&people)
	if err != nil {
		panic(err)
	}
	log.Println(people)
	return people, nil
}

func (o orderRepositoryImpl) UpdateOrderByID(ctx context.Context, id uint64, reqDataOrder *entity.Orders, reqDataItems entity.AllItems) error {
	queryItems := UPDATE_ITEM
	queryOrder := UPDATE_ORDER

	stmt, err := o.db.PrepareContext(ctx, queryOrder)
	if err != nil {
		log.Printf("[UpdateOrderByID] failed to prepare the statement, err => %v", err)
		return err
	}

	rows, err := stmt.ExecContext(
		ctx,
		reqDataOrder.CustomerName,
		id,
	)
	if err != nil {
		log.Printf("[UpdateOrderByID] failed to update the data, err => %v", err)
		return err
	}

	res1, _ := rows.RowsAffected()

	for _, items := range reqDataItems {
		stmt, err = o.db.PrepareContext(ctx, queryItems)
		if err != nil {
			log.Printf("[UpdateOrderByID] failed to prepare the statement, err => %v", err)
			return err
		}
		rows, err = stmt.Exec(
			items.Description,
			items.ItemCode,
			items.Quantity,
			items.ItemId,
		)
		if err != nil {
			log.Printf("[UpdateOrderByID] failed to update the data, err => %v", err)
			return err
		}

		res2, _ := rows.RowsAffected()
		res1 = res1 + res2
	}
	return nil
}

func (o orderRepositoryImpl) DeleteOrderByID(ctx context.Context, id uint64) (int, error) {
	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[DeleteOrderByID] failed to begin transaction, err => %v", err)
		return 0, err
	}
	defer tx.Rollback()

	queryDeleteItem := DELETE_ITEMS
	_, err = o.db.Query(queryDeleteItem, id)
	if err != nil {
		log.Printf("[DeleteByOrderID] failed to delete item data, err => %v", err)
		return 0, err
	}

	queryDeleteOrder := DELETE_ORDER
	_, err = o.db.Query(queryDeleteOrder, id)
	if err != nil {
		log.Printf("[DeleteOrderByID] failed to delete order data, err => %v", err)
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[DeleteOrderByID] transaction failed, err => %v", err)
		return 0, err
	}
	return int(id), nil
}

func (o orderRepositoryImpl) CheckOrders(ctx context.Context) (int, error) {
	query := COUNT_ORDERS

	stmt, err := o.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckOrders] failed to prepare the statement, err => %v", err)
		return 0, err
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[CheckOrders] failed to query to the database, err => %v", err)
		return 0, err
	}

	var orderCount int

	for rows.Next() {
		err := rows.Scan(
			&orderCount,
		)
		if err != nil {
			log.Printf("[CheckOrders] failed to scan data from database, err => %v", err)
			return 0, err
		}
	}
	return orderCount, nil
}

func (o orderRepositoryImpl) ViewAllOrders(ctx context.Context) (entity.OrdersJoined, error) {
	query := SELECT_ORDERS

	stmt, err := o.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewAllOrders] failed to prepare the statement, err => %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewAllOrders] failed to query to the database, err => %v", err)
		return nil, err
	}

	var ordersJoined entity.OrdersJoined

	for rows.Next() {
		orderItemJoined := entity.OrdersItemsJoined{}

		err := rows.Scan(
			&orderItemJoined.Orders.OrderID,
			&orderItemJoined.Orders.CustomerName,
			&orderItemJoined.Orders.CreatedAt,
			&orderItemJoined.Orders.UpdatedAt,
		)
		if err != nil {
			log.Printf("[ViewAllOrders] failed to scan data from database, err => %v", err)
			return nil, err
		}

		ordersJoined = append(ordersJoined, &orderItemJoined)
	}
	query = SELECT_ITEMS
	for idx, items := range ordersJoined {

		stmt, err = o.db.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("[ViewAllOrders] failed to prepare the statement, err => %v", err)
			return nil, err
		}
		rows, err := stmt.QueryContext(ctx, ordersJoined[idx].Orders.OrderID)
		if err != nil {
			log.Printf("[ViewAllOrders] failed to query to the database, err => %v", err)
			return nil, err
		}

		for rows.Next() {
			allItems := entity.Items{}

			err = rows.Scan(
				&allItems.ItemId,
				&allItems.ItemCode,
				&allItems.Description,
				&allItems.Quantity,
				&allItems.OrderID,
			)
			if err != nil {
				log.Printf("[ViewAllOrders] failed to scan data from database, err => %v", err)
				return nil, err
			}

			items.Items = append(items.Items, &allItems)
		}
	}
	return ordersJoined, nil
}

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
