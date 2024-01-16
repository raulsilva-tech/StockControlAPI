package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/configs"
	"github.com/raulsilva-tech/StockControlAPI/internal/infra/database"
	"github.com/raulsilva-tech/StockControlAPI/internal/infra/webserver/handlers"
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
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ptDAO := database.NewProductTypeDAO(db)
	ptHandler := handlers.NewProductTypeHandler(*ptDAO)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/product_types", func(r chi.Router) {
		r.Post("/", ptHandler.CreateProductType)
		r.Get("/", ptHandler.GetAllProductType)
		r.Get("/{id}", ptHandler.GetProductType)
		r.Put("/{id}", ptHandler.UpdateProductType)
		r.Delete("/{id}", ptHandler.DeleteProductType)
	})

	http.ListenAndServe(":8888", r)

}
