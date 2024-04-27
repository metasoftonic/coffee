package main

import (
	"fmt"
	"github/metasoftonic/coffee-api/db/postgres"
	"github/metasoftonic/coffee-api/internals/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbSource := os.Getenv("DATASOURCE")
	PORT := os.Getenv("PORT")
	fmt.Println(PORT)

	store, err := postgres.NewStore(dbSource)
	if err != nil {
		log.Fatal(err)
	}

	h := handlers.NewHandler(store)
	if err := http.ListenAndServe(PORT, h); err != nil {
		log.Fatal(err)
	}
}
