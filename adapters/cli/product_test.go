package cli_test

import (
	"fmt"
	"github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/emmanuelviniciusdev/imersao-fullcycle-arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productID := "be0dda1e-09de-4952-97ca-10c143a9d5ab"
	productName := "Limão"
	productPrice := 0.25
	productStatus := "ENABLED"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)

	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	result, err := cli.Run(productServiceMock, "create", productID, productName, productPrice)

	resultExpected := fmt.Sprintf("Product (%s) with name %s has been created with price %f and status %s",
		productID, productName, productPrice, productStatus)

	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	result, err = cli.Run(productServiceMock, "enable", productID, productName, productPrice)

	resultExpected = fmt.Sprintf("Product (%s) has been enabled", productID)

	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	result, err = cli.Run(productServiceMock, "disable", productID, productName, productPrice)

	resultExpected = fmt.Sprintf("Product (%s) has been disabled", productID)

	require.Equal(t, resultExpected, result)
	require.Nil(t, err)

	result, err = cli.Run(productServiceMock, "", productID, productName, productPrice)

	resultExpected = fmt.Sprintf("Product ID: %s, Product Name: %s, Product Price: %f, Product Status: %s",
		productID, productName, productPrice, productStatus)

	require.Equal(t, resultExpected, result)
	require.Nil(t, err)
}
