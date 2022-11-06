package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/milena-mognon/fc-ports-and-adapters/adapters/db"
	"github.com/milena-mognon/fc-ports-and-adapters/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDbAdapter := db2.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Exemple", 30)

	productService.Enable(product)

}
