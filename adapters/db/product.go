package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/milena-mognon/fc-ports-and-adapters/application"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Save(product ProductInterface) (ProductInterface, error)
