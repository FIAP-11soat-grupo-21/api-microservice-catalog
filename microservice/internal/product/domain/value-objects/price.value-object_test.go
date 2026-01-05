package value_objects

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPrice_Invalid(t *testing.T) {
	_, err := NewPrice(0)
	require.Error(t, err)
	_, err = NewPrice(-10)
	require.Error(t, err)
}

func TestNewPrice_Valid(t *testing.T) {
	p, err := NewPrice(10.5)
	require.NoError(t, err)
	require.Equal(t, 10.5, p.Value())
}
