package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/router"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/global/server"
	"github.com/maheswaradevo/hacktiv8-assignment2/pkg/database"
)

func main() {
	log.Println("[main] starting server")
	config.Init()
	cfg := config.GetConfig()
	root := mux.NewRouter()
	db := database.GetDatabase()

	router.Init(root, db)

	s := server.ProvideServer(cfg.ServerAddress, root)
	s.ListenAndServe()
}
