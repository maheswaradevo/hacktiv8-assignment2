package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/maheswaradevo/hacktiv8-assignment2/docs"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/server"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/database"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Orders API
// @version 1.0
// @description This is service to managing order
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/license/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /

func main() {
	log.Println("[main] starting server")
	config.Init()
	cfg := config.GetConfig()
	root := mux.NewRouter()
	db := database.GetDatabase()

	router.Init(root, db)
	root.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	s := server.ProvideServer(cfg.ServerAddress, root)
	s.ListenAndServe()
}
