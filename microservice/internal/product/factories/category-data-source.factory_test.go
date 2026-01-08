package factories

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategoryDataSource_ReturnsICategoryDataSource(t *testing.T) {
	ds := NewCategoryDataSource()
	require.NotNil(t, ds)
}
