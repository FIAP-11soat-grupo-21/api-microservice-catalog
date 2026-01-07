package factories

import (
	"tech_challenge/internal/product/interfaces"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProductDataSource_ReturnsIProductDataSource(t *testing.T) {
	ds := NewProductDataSource()
	require.NotNil(t, ds)
	_, ok := ds.(interfaces.IProductDataSource)
	require.True(t, ok)
}
