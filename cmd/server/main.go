<<<<<<< HEAD
package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/raulsilva-tech/StockControlAPI/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/raulsilva-tech/StockControlAPI/configs"
	"github.com/raulsilva-tech/StockControlAPI/internal/infra/database"
	"github.com/raulsilva-tech/StockControlAPI/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title	Stock Control API
// @version	1.0
// @description Stock Control API
// @termsOfService	http://swagger.io/terms

// @contact.name	Raul Paes Silva
// @contact.url	http://github.com/raulsilva-tech
// @contact.email raulpaes.work@gmail.com

// @host	localhost:8888
// @BasePath /

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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	createRoutes(r, db)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8888/docs/doc.json")))

	fmt.Println("Server started on port " + cfg.WebServerPort)
	http.ListenAndServe(":"+cfg.WebServerPort, r)

}

func createRoutes(r *chi.Mux, db *sql.DB) {

	//product type
	ptDAO := database.NewProductTypeDAO(db)
	ptHandler := handlers.NewProductTypeHandler(*ptDAO)
	r.Route("/product_types", func(r chi.Router) {
		r.Post("/", ptHandler.CreateProductType)
		r.Get("/", ptHandler.GetAllProductType)
		r.Get("/{id}", ptHandler.GetProductType)
		r.Put("/{id}", ptHandler.UpdateProductType)
		r.Delete("/{id}", ptHandler.DeleteProductType)
	})

	//product
	pDAO := database.NewProductDAO(db)
	pHandler := handlers.NewProductHandler(*pDAO)
	r.Route("/products", func(r chi.Router) {
		r.Post("/", pHandler.CreateProduct)
		r.Get("/", pHandler.GetAllProduct)
		r.Get("/{id}", pHandler.GetProduct)
		r.Put("/{id}", pHandler.UpdateProduct)
		r.Delete("/{id}", pHandler.DeleteProduct)
	})

	//label
	lDAO := database.NewLabelDAO(db)
	lHandler := handlers.NewLabelHandler(*lDAO)
	r.Route("/labels", func(r chi.Router) {
		r.Post("/", lHandler.CreateLabel)
		r.Get("/", lHandler.GetAllLabel)
		r.Get("/{id}", lHandler.GetLabel)
		r.Put("/{id}", lHandler.UpdateLabel)
		r.Delete("/{id}", lHandler.DeleteLabel)
	})

	//operation
	opDAO := database.NewOperationDAO(db)
	opHandler := handlers.NewOperationHandler(*opDAO)
	r.Route("/operations", func(r chi.Router) {
		r.Post("/", opHandler.CreateOperation)
		r.Get("/", opHandler.GetAllOperation)
		r.Get("/{id}", opHandler.GetOperation)
		r.Put("/{id}", opHandler.UpdateOperation)
		r.Delete("/{id}", opHandler.DeleteOperation)
	})

	//stock
	stDAO := database.NewStockDAO(db)
	stHandler := handlers.NewStockHandler(*stDAO)
	r.Route("/stocks", func(r chi.Router) {
		r.Post("/", stHandler.CreateStock)
		r.Get("/", stHandler.GetAllStock)
		r.Get("/{id}", stHandler.GetStock)
		r.Put("/{id}", stHandler.UpdateStock)
		r.Delete("/{id}", stHandler.DeleteStock)
	})

	//stock product
	spDAO := database.NewStockProductDAO(db)
	spHandler := handlers.NewStockProductHandler(*spDAO)
	r.Route("/stock_products", func(r chi.Router) {
		r.Post("/", spHandler.CreateStockProduct)
		r.Get("/", spHandler.GetAllStockProduct)
		r.Get("/{id}", spHandler.GetStockProduct)
		r.Put("/{id}", spHandler.UpdateStockProduct)
		r.Delete("/{id}", spHandler.DeleteStockProduct)
	})

	//transaction
	trDAO := database.NewTransactionDAO(db)
	trHandler := handlers.NewTransactionHandler(*trDAO)
	r.Route("/transactions", func(r chi.Router) {
		r.Post("/", trHandler.CreateTransaction)
		r.Get("/", trHandler.GetAllTransaction)
		r.Get("/{id}", trHandler.GetTransaction)
		r.Put("/{id}", trHandler.UpdateTransaction)
		r.Delete("/{id}", trHandler.DeleteTransaction)
	})

	//user
	userDAO := database.NewUserDAO(db)
	userHandler := handlers.NewUserHandler(*userDAO)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetAllUser)
		r.Get("/{id}", userHandler.GetUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
		r.Post("/login", userHandler.Login)
		r.Get("/logout/{id}", userHandler.Logout)
	})

	//User session
	sessionDAO := database.NewUserSessionDAO(db)
	sessionHandler := handlers.NewUserSessionHandler(*sessionDAO)
	r.Route("/user_sessions", func(r chi.Router) {
		r.Post("/", sessionHandler.CreateUserSession)
		r.Get("/", sessionHandler.GetAllUserSession)
		r.Get("/{id}", sessionHandler.GetUserSession)
		r.Put("/{id}", sessionHandler.UpdateUserSession)
		r.Delete("/{id}", sessionHandler.DeleteUserSession)
	})

	//user operations
	uoDAO := database.NewUserOperationDAO(db)
	uoHandler := handlers.NewUserOperationHandler(*uoDAO)
	r.Route("/user_operations", func(r chi.Router) {
		r.Post("/", uoHandler.CreateUserOperation)
		r.Get("/", uoHandler.GetAllUserOperation)
		r.Get("/{id}", uoHandler.GetUserOperation)
		r.Put("/{id}", uoHandler.UpdateUserOperation)
		r.Delete("/{id}", uoHandler.DeleteUserOperation)
	})
}
=======
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

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	createRoutes(r, db)

	http.ListenAndServe(":8888", r)

}

func createRoutes(r *chi.Mux, db *sql.DB) {

	//product type
	ptDAO := database.NewProductTypeDAO(db)
	ptHandler := handlers.NewProductTypeHandler(*ptDAO)
	r.Route("/product_types", func(r chi.Router) {
		r.Post("/", ptHandler.CreateProductType)
		r.Get("/", ptHandler.GetAllProductType)
		r.Get("/{id}", ptHandler.GetProductType)
		r.Put("/{id}", ptHandler.UpdateProductType)
		r.Delete("/{id}", ptHandler.DeleteProductType)
	})

	//product
	pDAO := database.NewProductDAO(db)
	pHandler := handlers.NewProductHandler(*pDAO)
	r.Route("/products", func(r chi.Router) {
		r.Post("/", pHandler.CreateProduct)
		r.Get("/", pHandler.GetAllProduct)
		r.Get("/{id}", pHandler.GetProduct)
		r.Put("/{id}", pHandler.UpdateProduct)
		r.Delete("/{id}", pHandler.DeleteProduct)
	})

	//label
	lDAO := database.NewLabelDAO(db)
	lHandler := handlers.NewLabelHandler(*lDAO)
	r.Route("/labels", func(r chi.Router) {
		r.Post("/", lHandler.CreateLabel)
		r.Get("/", lHandler.GetAllLabel)
		r.Get("/{id}", lHandler.GetLabel)
		r.Put("/{id}", lHandler.UpdateLabel)
		r.Delete("/{id}", lHandler.DeleteLabel)
	})

	//operation
	opDAO := database.NewOperationDAO(db)
	opHandler := handlers.NewOperationHandler(*opDAO)
	r.Route("/operations", func(r chi.Router) {
		r.Post("/", opHandler.CreateOperation)
		r.Get("/", opHandler.GetAllOperation)
		r.Get("/{id}", opHandler.GetOperation)
		r.Put("/{id}", opHandler.UpdateOperation)
		r.Delete("/{id}", opHandler.DeleteOperation)
	})

	//stock
	stDAO := database.NewStockDAO(db)
	stHandler := handlers.NewStockHandler(*stDAO)
	r.Route("/stocks", func(r chi.Router) {
		r.Post("/", stHandler.CreateStock)
		r.Get("/", stHandler.GetAllStock)
		r.Get("/{id}", stHandler.GetStock)
		r.Put("/{id}", stHandler.UpdateStock)
		r.Delete("/{id}", stHandler.DeleteStock)
	})

	//stock product
	spDAO := database.NewStockProductDAO(db)
	spHandler := handlers.NewStockProductHandler(*spDAO)
	r.Route("/stock_products", func(r chi.Router) {
		r.Post("/", spHandler.CreateStockProduct)
		r.Get("/", spHandler.GetAllStockProduct)
		r.Get("/{id}", spHandler.GetStockProduct)
		r.Put("/{id}", spHandler.UpdateStockProduct)
		r.Delete("/{id}", spHandler.DeleteStockProduct)
	})

	//transaction
	trDAO := database.NewTransactionDAO(db)
	trHandler := handlers.NewTransactionHandler(*trDAO)
	r.Route("/transactions", func(r chi.Router) {
		r.Post("/", trHandler.CreateTransaction)
		r.Get("/", trHandler.GetAllTransaction)
		r.Get("/{id}", trHandler.GetTransaction)
		r.Put("/{id}", trHandler.UpdateTransaction)
		r.Delete("/{id}", trHandler.DeleteTransaction)
	})

	//user
	userDAO := database.NewUserDAO(db)
	userHandler := handlers.NewUserHandler(*userDAO)
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Get("/", userHandler.GetAllUser)
		r.Get("/{id}", userHandler.GetUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
		r.Post("/login", userHandler.Login)
		r.Get("/logout/{id}", userHandler.Logout)
	})

	//User session
	sessionDAO := database.NewUserSessionDAO(db)
	sessionHandler := handlers.NewUserSessionHandler(*sessionDAO)
	r.Route("/user_sessions", func(r chi.Router) {
		r.Post("/", sessionHandler.CreateUserSession)
		r.Get("/", sessionHandler.GetAllUserSession)
		r.Get("/{id}", sessionHandler.GetUserSession)
		r.Put("/{id}", sessionHandler.UpdateUserSession)
		r.Delete("/{id}", sessionHandler.DeleteUserSession)
	})

	//user operations
	uoDAO := database.NewUserOperationDAO(db)
	uoHandler := handlers.NewUserOperationHandler(*uoDAO)
	r.Route("/user_operations", func(r chi.Router) {
		r.Post("/", uoHandler.CreateUserOperation)
		r.Get("/", uoHandler.GetAllUserOperation)
		r.Get("/{id}", uoHandler.GetUserOperation)
		r.Put("/{id}", uoHandler.UpdateUserOperation)
		r.Delete("/{id}", uoHandler.DeleteUserOperation)
	})
}
>>>>>>> d4eba3be9444a00975090f26358cb6323f2e2548
