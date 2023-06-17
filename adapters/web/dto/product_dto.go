package dto

import "github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"

type ProductDTO struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProductDTO(ID string, name string, price float64, status string) *ProductDTO {
	return &ProductDTO{ID: ID, Name: name, Price: price, Status: status}
}

func (p *ProductDTO) Bind(product *application.Product) (*application.Product, error) {
	if p.ID != "" {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status

	_, err := product.IsValid()

	if err != nil {
		return &application.Product{}, err
	}

	return product, nil
}
