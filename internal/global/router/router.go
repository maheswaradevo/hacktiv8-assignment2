package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/ping"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/order"
)

func Init(router *mux.Router, db *sql.DB) {
	webRouter := router.NewRoute().PathPrefix("/api/v1").Subrouter()

	pingService := ping.NewPingService()
	pingController := ping.NewPingHandler(webRouter, pingService)
	pingController.InitHandler()

	orderService := order.ProvideOrderService(db)
	orderController := order.ProvideOrderHandler(webRouter, orderService)
	orderController.InitHandler()
}
