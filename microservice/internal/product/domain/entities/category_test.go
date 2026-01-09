package entities

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategory_InvalidName(t *testing.T) {
	_, err := NewCategory("id", "a", true)
	require.Error(t, err)
}

func TestNewCategory_Valid(t *testing.T) {
	id := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	c, err := NewCategory(id, "Bebidas", true)
	require.NoError(t, err)
	require.Equal(t, id, c.ID)
	require.Equal(t, "Bebidas", c.Name.Value())
	require.True(t, c.Active)
}

func TestCategory_SetName_Invalid(t *testing.T) {
	c, err := NewCategory("id", "Bebidas", true)
	require.NoError(t, err)
	require.Error(t, c.SetName("a"))
}

func TestCategory_SetName_Valid(t *testing.T) {
	c, err := NewCategory("id", "Bebidas", true)
	require.NoError(t, err)
	require.NoError(t, c.SetName("Refrigerantes"))
	require.Equal(t, "Refrigerantes", c.Name.Value())
}
