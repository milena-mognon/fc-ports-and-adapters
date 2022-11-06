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

func (p *ProductDB) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int
	p.db.QueryRow("Select id from products where id=?", product.GetID()).Scan(&rows)

	if rows == 0 {
		_, err := p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}
	return product, nil
}

func (p *ProductDB) create(product application.ProductInterface) (application.ProductInterface, error) { // método não exportado
	stmt, err := p.db.Prepare(`insert into products(id, name, price, status) values (?,?,?,?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		product.GetID(),
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
	)

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product application.ProductInterface) (application.ProductInterface, error) { // método não exportado
	_, err := p.db.Exec("update products set name=?, price=?, status=? where id=?",
		product.GetName(),
		product.GetPrice(),
		product.GetStatus(),
		product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
