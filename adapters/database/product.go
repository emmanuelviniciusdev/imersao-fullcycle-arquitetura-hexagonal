package database

import (
	"database/sql"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	db *sql.DB
}

func NewProductDatabase(db *sql.DB) *ProductDatabase {
	return &ProductDatabase{db: db}
}

func (p *ProductDatabase) Get(ID string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM tb_product WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(ID).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDatabase) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var totalProductsFound int

	stmt, err := p.db.Prepare("SELECT COUNT(*) FROM tb_product WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(product.GetID()).Scan(&totalProductsFound)

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	if totalProductsFound == 0 {
		_, err = p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err = p.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDatabase) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE tb_product SET name = ?, price = ?, status = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDatabase) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO tb_product (id, name, price, status) VALUES (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}
