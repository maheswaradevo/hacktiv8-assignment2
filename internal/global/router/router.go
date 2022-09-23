package router

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/ping"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/auth"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/order"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/middleware"
)

func Init(router *mux.Router, db *sql.DB) {
	webRouter := router.NewRoute().PathPrefix("/api/v1").Subrouter()
	protectedRoute := router.NewRoute().PathPrefix("/api/v1").Subrouter()

	pingService := ping.NewPingService()
	pingController := ping.NewPingHandler(webRouter, pingService)
	pingController.InitHandler()

	protectedRoute.Use(middleware.AuthMiddleware())

	orderService := order.ProvideOrderService(db)
	orderController := order.ProvideOrderHandler(webRouter, protectedRoute, orderService)
	orderController.InitHandler()

	authService := auth.ProvideAuthService(db)
	authController := auth.ProvideAuthHandler(webRouter, authService)
	authController.InitHandler()
}
