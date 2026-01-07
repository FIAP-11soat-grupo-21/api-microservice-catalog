package daos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductImageDAO_TableName(t *testing.T) {
	dao := ProductImageDAO{}
	require.Equal(t, "product_images", dao.TableName())
}
