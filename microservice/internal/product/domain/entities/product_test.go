package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProduct_InvalidName(t *testing.T) {
	_, err := NewProduct("id", "", "catid", "desc", 10.0, true)
	require.Error(t, err)
}

func TestNewProduct_Valid(t *testing.T) {
	id := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	catid := "b3bb189e-8bf9-3888-9912-ace4e6543002"
	desc := "Refrigerante"
	p, err := NewProduct(id, "Coca-Cola", catid, desc, 5.99, true)
	require.NoError(t, err)
	require.Equal(t, id, p.ID)
	require.Equal(t, "Coca-Cola", p.Name.Value())
	require.Equal(t, catid, p.CategoryID)
	require.Equal(t, desc, p.Description)
	require.Equal(t, 5.99, p.Price)
	require.True(t, p.Active)
}

func TestProduct_SetName_Invalid(t *testing.T) {
	p, err := NewProduct("id", "Coca-Cola", "catid", "desc", 5.99, true)
	require.NoError(t, err)
	require.Error(t, p.SetName(""))
}

func TestProduct_SetName_Valid(t *testing.T) {
	p, err := NewProduct("id", "Coca-Cola", "catid", "desc", 5.99, true)
	require.NoError(t, err)
	require.NoError(t, p.SetName("Guaraná"))
	require.Equal(t, "Guaraná", p.Name.Value())
}
