package main

import (
	"mezink-goland-assessment/config"
	"mezink-goland-assessment/handlers"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.NewConfig()

	db, err := sqlx.Open("postgres", cfg.DBConnectionString())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	recordHandler := &handlers.RecordHandler{DB: db}

	http.HandleFunc("/api/records", recordHandler.InsertAndFetch)
	http.HandleFunc("/api/fetch-table", recordHandler.FetchTable)

	http.ListenAndServe(":8080", nil)
}
