package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResponseJson(t *testing.T) {
	message := "Hello, JSON"

	responseJson := ResponseJson(message)
	responseJsonExpected := []byte(`{"message":"Hello, JSON"}`)

	require.Equal(t, responseJsonExpected, responseJson)
}
