package application_test

import (
	"testing"

	"github.com/renatospaka/fc2.0-ddd/application"
	"github.com/stretchr/testify/require"
	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Price:  10,
		Status: application.DISABLED,
	}

	//execute the enable method with a valid price and test the result. Must work
	err := product.Enable()
	require.Nil(t, err)

	//execute the enable method with an invalid price and test the result. Must fail
	product.Price = 0
	err = product.Enable()
	require.Equal(t, "Price must be grater than zero", err.Error())
} 

func TestProduct_Disable (t *testing.T) {
	product := application.Product{
		Name:   "Hello",
		Price:  0,
		Status: application.ENABLED,
	}

	//execute the disable method with a valid price and test the result. Must work
	err := product.Disable()
	require.Nil(t, err)
	
	//execute the disable method with an invalid price and test the result. Must fail
	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Price must be equal to zero", err.Error())
}

func TestProduct_IsValid (t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Hello",
		Price:  10,
		Status: application.DISABLED,
	}

	_, err := product.IsValid()
	require.Nil(t, err)
	//require.BoolAssertionFunc(t, true, valid)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "Status must be enabled or disabled", err.Error())
	
	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "Price must be greater or equal to zero", err.Error())
}