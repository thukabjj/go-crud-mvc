package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/thukabjj/go-crud-mvc/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	routes.CarregaRotas()

	http.ListenAndServe(":8080", nil)
}
