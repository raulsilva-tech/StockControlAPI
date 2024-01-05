package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/configs"
)

type Product struct {
	Id          string
	Description string
}

func main() {

	//load configuration
	cfg, _ := configs.LoadConfig(".")

	//starting database connection
	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	fmt.Println(DataSourceName)
	db, err := sql.Open(cfg.DBDriver, DataSourceName)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// selectAllProducts(db)
}

// func selectAllProducts(db *sql.DB) error {
// 	rows, err := db.Query("select id,description from products")
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	var products []Product

// 	for rows.Next() {
// 		var p Product
// 		err = rows.Scan(&p.Id, &p.Description)
// 		products = append(products, p)
// 		fmt.Println(p)
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
