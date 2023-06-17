package cli

import (
	"fmt"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
)

func Run(
	service application.ProductServiceInterface, action string, productID string,
	productName string, productPrice float64) (string, error) {
	result := ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product (%s) with name %s has been created with price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productID)

		if err != nil {
			return "", err
		}

		_, err = service.Enable(product)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product (%s) has been enabled", productID)
	case "disable":
		product, err := service.Get(productID)

		if err != nil {
			return "", err
		}

		_, err = service.Disable(product)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product (%s) has been disabled", productID)
	default:
		product, err := service.Get(productID)

		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("Product ID: %s, Product Name: %s, Product Price: %f, Product Status: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
