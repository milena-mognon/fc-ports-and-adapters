package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/milena-mognon/fc-ports-and-adapters/application"
	mock_application "github.com/milena-mognon/fc-ports-and-adapters/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish() // espera tudo o que está acontencedo no método para depois executar o comando

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	/**
		* Espero que todas as vezes que o método Get for chamado, independente do valor que passar no id,
		* quero que retorne um product e um nil, independente de quantas vezes for chamado
	  *
		* Simula que aquando chama o ,étodo Get retorna um produto
	*/
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")

	require.Nil(t, err)

	require.Equal(t, product, result)
}

func TestProductService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish() // espera tudo o que está acontencedo no método para depois executar o comando

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)

	require.Nil(t, err)

	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish() // espera tudo o que está acontencedo no método para depois executar o comando

	product := mock_application.NewMockProductInterface(ctrl)

	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)

	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}
