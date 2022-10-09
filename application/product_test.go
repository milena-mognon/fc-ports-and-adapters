package application_test

import (
	"testing"

	"github.com/milena-mognon/fc-ports-and-adapters/application"
	"github.com/stretchr/testify/require"
)

// se quiser testar as funções/métodos internos do pacote usa o package com mesmo nome
// se quiser testar de forma externa coloca o nome do pacote com _teste

// Exemplo: método Enable -> permite externo
// Exemplo: método enable = somente interno

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()

	require.Equal(t, "the price must be zero in order to disable the product", err.Error())
}
