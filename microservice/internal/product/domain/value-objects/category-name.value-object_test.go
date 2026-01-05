package value_objects

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategoryName_InvalidShort(t *testing.T) {
	_, err := NewCategoryName("ab")
	require.Error(t, err)
}

func TestNewCategoryName_InvalidLong(t *testing.T) {
	longName := strings.Repeat("a", 101)
	_, err := NewCategoryName(longName)
	require.Error(t, err)
}

func TestNewCategoryName_Valid(t *testing.T) {
	name, err := NewCategoryName("Bebidas")
	require.NoError(t, err)
	require.Equal(t, "Bebidas", name.Value())
}
