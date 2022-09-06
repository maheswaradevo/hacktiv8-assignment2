package main

import (
	"hacktiv8-assignment2/pkg/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.GetDatabase(":3306", "root", "pundadevo25052002", "testing")
}
