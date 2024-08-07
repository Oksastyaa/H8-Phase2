package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"ngc_release2/db"
	"os"
	"time"
)

type application struct {
	port     int
	infoLog  *log.Logger
	errorLog *log.Logger
	db       *sql.DB
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	app.infoLog.Printf("Starting server on port %d", app.port)

	return srv.ListenAndServe()
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect db
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer database.Close()

	// Initialize application with the database connection
	app := &application{
		port:     8080,
		infoLog:  infoLog,
		errorLog: errorLog,
		db:       database,
	}
	// Start the server
	err = app.serve()
	if err != nil {
		app.errorLog.Fatal(err)
	}
}
