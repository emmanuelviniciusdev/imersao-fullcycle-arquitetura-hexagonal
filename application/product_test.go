package application_test

import (
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}

	product.ID = uuid.NewV4().String()
	product.Name = "Limão"
	product.Price = 0
	product.Status = application.DISABLED

	isValid, err := product.IsValid()

	require.Equal(t, true, isValid)
	require.Nil(t, err)

	product.Status = "FOO"

	isValid, err = product.IsValid()

	require.Equal(t, false, isValid)
	require.Equal(t, "the status must be ENABLED or DISABLED", err.Error())

	product.Status = application.DISABLED
	product.Price = -10

	isValid, err = product.IsValid()

	require.Equal(t, false, isValid)
	require.Equal(t, "the price cannot be negative", err.Error())

	product.Price = 0
	product.ID = "INVALID-UUID-V4"

	isValid, err = product.IsValid()

	require.Equal(t, false, isValid)
	require.Equal(t, "ID: INVALID-UUID-V4 does not validate as uuidv4", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Limão"
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 1.25

	err = product.Disable()
	require.Equal(t, "the price must be equal 0 in order to disable the product", err.Error())
}

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Price = 10.5

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than 0 in order to enable the product", err.Error())
}
