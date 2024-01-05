package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	Id    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		Id:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// product := NewProduct("Laranjas", 2.09)
	// err = insertProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(product)

	// product.Name = "Cenoura"
	// err = updateProduct(db, product)
	// if err != nil {
	// 	panic(err)
	// }

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Millisecond*500)
	// p, err := selectProduct(ctx, db, product.Id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)

	products, err := selectAllProducts(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

	// excluindo produto
	err = deleteProduct(ctx, db, products[len(products)-1].Id)
	if err != nil {
		panic(err)
	}
	products, err = selectAllProducts(ctx, db)
	if err != nil {
		panic(err)
	}
	for _, product := range products {
		fmt.Println(product)
	}

}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id,name,price) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Id, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name=?, price=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Id)
	if err != nil {
		return err
	}
	return nil

}

func selectProduct(ctx context.Context, db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	// err = stmt.QueryRow(id).Scan(&p.Id, &p.Name, &p.Price)
	err = stmt.QueryRowContext(ctx, id).Scan(&p.Id, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}
	return &p, nil

}

func deleteProduct(ctx context.Context, db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}
	return nil

}

func selectAllProducts(ctx context.Context, db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id,name,price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.Id, &p.Name, &p.Price)
		products = append(products, p)
	}

	if err != nil {
		return nil, err
	}
	return products, nil

}
