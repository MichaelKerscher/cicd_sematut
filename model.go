// model.go

package main

import (
	"database/sql"
)

type product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price, category FROM products WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price, &p.Category)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2, category=$3 WHERE id=$4",
		p.Name, p.Price, p.Category, p.ID)
	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO products(name, price, category) VALUES($1, $2, $3) RETURNING id",
		p.Name, p.Price, p.Category).Scan(&p.ID)
}

func getProducts(db *sql.DB, start, count int, name, sort string) ([]product, error) {
	order := "ASC"
	if sort == "desc" {
		order = "DESC"
	}

	query := "SELECT id, name, price, category FROM products WHERE name ILIKE $1 ORDER BY price " + order + " LIMIT $2 OFFSET $3"
	rows, err := db.Query(query, "%"+name+"%", count, start)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []product{}
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Category); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
