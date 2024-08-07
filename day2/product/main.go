package main

import (
	"Phase2/day2/product/handler"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	router := httprouter.New()
	router.GET("/products", handler.GetProducts(db))
	router.GET("/product/:id", handler.GetProduct(db))
	router.POST("/product", handler.CreateProduct(db))
	router.PUT("/product/:id", handler.UpdateProduct(db))
	//	listen and serve on
	log.Fatal(http.ListenAndServe(":8080", router))
}
