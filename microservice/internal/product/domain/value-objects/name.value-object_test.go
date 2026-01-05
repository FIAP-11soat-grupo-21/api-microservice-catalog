package value_objects

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewName_InvalidShort(t *testing.T) {
	_, err := NewName("a")
	require.Error(t, err)
}

func TestNewName_InvalidLong(t *testing.T) {
	longName := "a" + strings.Repeat("b", 100)
	_, err := NewName(longName)
	require.Error(t, err)
}

func TestNewName_Valid(t *testing.T) {
	name, err := NewName("Coca-Cola")
	require.NoError(t, err)
	require.Equal(t, "Coca-Cola", name.Value())
}
