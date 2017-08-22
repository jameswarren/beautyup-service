package main

import (
    "database/sql"
    //"log"
)

type product struct {
    ID    int     `json:"id"`
    Brand  string  `json:"brand"`
    Name string `json:"name"`
}

// func (p *product) getProduct(db *sql.DB) error {
//     return db.QueryRow("SELECT name, brand FROM products WHERE id=$1",
//         p.ID).Scan(&p.Name, &p.Brand)
// }

// func (p *product) updateProduct(db *sql.DB) error {
//     _, err :=
//         db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
//             p.Name, p.Price, p.ID)
//
//     return err
// }
//
// func (p *product) deleteProduct(db *sql.DB) error {
//     _, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)
//
//     return err
// }
//
// func (p *product) createProduct(db *sql.DB) error {
//     err := db.QueryRow(
//         "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
//         p.Name, p.Price).Scan(&p.ID)
//
//     if err != nil {
//         return err
//     }
//
//     return nil
// }

func getProducts(db *sql.DB, pattern string, count int) ([]product, error) {
    rows, err := db.Query(
        "SELECT id, brand, name FROM products WHERE (name LIKE '%' || $1 || '%') OR (brand LIKE '%' || $1 || '%') LIMIT $2",
        pattern, count)

    if err != nil {
    //  log.Fatal("pattern=%s count=%d", pattern, count)
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
