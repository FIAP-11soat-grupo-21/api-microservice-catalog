package factories

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProductDataSource_ReturnsIProductDataSource(t *testing.T) {
	ds := NewProductDataSource()
	require.NotNil(t, ds)
}
