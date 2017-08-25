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


func getProducts(db *sql.DB, pattern string, count int) ([]product, error) {
    rows, err := db.Query(
        "SELECT id, brand, name FROM products WHERE (name ILIKE '%' || $1 || '%') OR (brand ILIKE '%' || $1 || '%') LIMIT $2",
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
