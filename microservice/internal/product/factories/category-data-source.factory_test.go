package factories

import (
	"tech_challenge/internal/product/interfaces"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategoryDataSource_ReturnsICategoryDataSource(t *testing.T) {
	ds := NewCategoryDataSource()
	require.NotNil(t, ds)
	_, ok := ds.(interfaces.ICategoryDataSource)
	require.True(t, ok)
}
