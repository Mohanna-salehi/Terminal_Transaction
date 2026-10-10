package main

import (
	"Terminal_Transaction/internal/api"
	"Terminal_Transaction/internal/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	h := api.NewHandler(db)

	ip := http.NewServeMux()

	ip.HandleFunc("GET /api/reports/terminals", h.GetTerminals)

	ip.HandleFunc("GET /api/reports/terminals/daily", h.GetDailyTransaction)

	ip.HandleFunc("GET /api/reports/terminals/last-transactions", h.GetLastTransaction)

	err = http.ListenAndServe(":8080", ip)
	if err != nil {
		log.Fatal(err)
	}
}
