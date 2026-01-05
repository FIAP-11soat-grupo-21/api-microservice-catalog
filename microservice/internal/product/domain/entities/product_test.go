package entities

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	_ = godotenv.Load("../../../.env.local.example")
	os.Setenv("GO_ENV", "test")
	os.Setenv("API_PORT", "8080")
	os.Setenv("API_HOST", "localhost")
	os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
	os.Setenv("DB_RUN_MIGRATIONS", "false")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_NAME", "test_db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "test_user")
	os.Setenv("DB_PASSWORD", "test_pass")
	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
	os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
	code := m.Run()
	os.Exit(code)
}

func TestNewProduct_EmptyName(t *testing.T) {
	_, err := NewProduct("id", "", "catid", "desc", 10.0, true)
	require.NoError(t, err)
}

func TestNewProduct_Valid(t *testing.T) {
	id := "a3bb189e-8bf9-3888-9912-ace4e6543002"
	catid := "b3bb189e-8bf9-3888-9912-ace4e6543002"
	desc := "Refrigerante"
	p, err := NewProduct(id, catid, "Coca-Cola", desc, 5.99, true)
	require.NoError(t, err)
	require.Equal(t, id, p.ID)
	require.Equal(t, catid, p.CategoryID)
	require.Equal(t, desc, p.Description)
	require.Equal(t, 5.99, p.Price.Value())
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
