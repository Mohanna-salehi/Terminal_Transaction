package main

import (
	"Terminal_Transaction/internal/database"
	"Terminal_Transaction/internal/seed"
	"context"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = seed.Seed(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")

}
