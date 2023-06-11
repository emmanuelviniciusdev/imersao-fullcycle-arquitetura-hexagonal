package main

import (
	"database/sql"
	"fmt"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/adapters/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "database.sqlite3")

	productDatabaseAdapter := database.NewProductDatabase(db)

	productService := application.NewProductService(productDatabaseAdapter)

	product, err := productService.Create("Pumpkin", 0.75)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Product:", product)
}
