package database_test

import (
	"database/sql"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/adapters/database"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var db *sql.DB

func setUp() {
	db, _ = sql.Open("sqlite3", ":memory:")

	createTables(db)
	insertProducts(db)
}

func setDown() {
	db.Close()
}

func executeQueries(queries []string) {
	for _, query := range queries {
		stmt, err := db.Prepare(query)

		if err != nil {
			log.Fatal(err.Error())
		}

		_, err = stmt.Exec()

		if err != nil {
			log.Fatal(err.Error())
		}

		err = stmt.Close()

		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func createTables(db *sql.DB) {
	queries := []string{
		"CREATE TABLE tb_product (id STRING, name STRING, price FLOAT, status STRING)",
	}

	executeQueries(queries)
}

func insertProducts(db *sql.DB) {
	queries := []string{
		"INSERT INTO tb_product VALUES ('be0dda1e-09de-4952-97ca-10c143a9d5ab', 'Limão', 0.25, 'ENABLED')",
		"INSERT INTO tb_product VALUES ('2ff514dd-7b4d-4cf8-ac1a-8d1b8a2626ff', 'Cebola', 0, 'DISABLED')",
		"INSERT INTO tb_product VALUES ('0a19f70b-c9bf-4c90-a43f-bfbfcda6c8ee', 'Beterraba', 0.55, 'ENABLED')",
	}

	executeQueries(queries)
}

func TestProductDatabase_Save(t *testing.T) {
	setUp()
	defer setDown()

	productDatabase := database.NewProductDatabase(db)

	product := application.NewProduct()

	product.Name = "Pumpkin"
	product.Price = 0.75
	product.Enable()

	foundProduct, err := productDatabase.Get(product.GetID())
	require.Nil(t, foundProduct)
	require.NotNil(t, err)

	_, err = productDatabase.Save(product)
	require.Nil(t, err)

	foundProduct, err = productDatabase.Get(product.GetID())
	require.Equal(t, foundProduct, product)
	require.Nil(t, err)

	product.Price = 0
	product.Disable()

	_, err = productDatabase.Save(product)
	require.Nil(t, err)

	foundProduct, err = productDatabase.Get(product.GetID())
	require.Equal(t, foundProduct, product)
	require.Nil(t, err)
}

func TestProductDatabase_Get(t *testing.T) {
	setUp()
	defer setDown()

	productDatabase := database.NewProductDatabase(db)

	product, err := productDatabase.Get("0a19f70b-c9bf-4c90-a43f-bfbfcda6c8ee")

	require.Nil(t, err)
	require.Equal(t, "0a19f70b-c9bf-4c90-a43f-bfbfcda6c8ee", product.GetID())
	require.Equal(t, "Beterraba", product.GetName())
	require.Equal(t, 0.55, product.GetPrice())
	require.Equal(t, "ENABLED", product.GetStatus())
}
