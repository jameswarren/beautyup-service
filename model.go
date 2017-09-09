package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Brand string  `json:"brand"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Image string  `json:"image"`
}

func getProducts(db *sql.DB, pattern string, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, brand, name, price, image_link FROM products WHERE (name ILIKE '%' || $1 || '%') OR (brand ILIKE '%' || $1 || '%') LIMIT $2",
		pattern, count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Brand, &p.Name); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
