package entities

import (
	"os"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	testenv "tech_challenge/internal/shared/test"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// func TestMain(m *testing.M) {
// 	_ = godotenv.Load("../../../.env.local.example")
// 	os.Setenv("GO_ENV", "test")
// 	os.Setenv("API_PORT", "8080")
// 	os.Setenv("API_HOST", "localhost")
// 	os.Setenv("API_UPLOAD_URL", "http://localhost:8080/uploads")
// 	os.Setenv("DB_RUN_MIGRATIONS", "false")
// 	os.Setenv("DB_HOST", "localhost")
// 	os.Setenv("DB_NAME", "test_db")
// 	os.Setenv("DB_PORT", "5432")
// 	os.Setenv("DB_USERNAME", "test_user")
// 	os.Setenv("DB_PASSWORD", "test_pass")
// 	os.Setenv("AWS_REGION", "us-east-1")
// 	os.Setenv("AWS_S3_BUCKET_NAME", "test-bucket")
// 	os.Setenv("AWS_S3_PRESIGN_EXPIRATION", "3600")
// 	code := m.Run()
// 	os.Exit(code)
// }

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
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

func TestNewProduct_InvalidName(t *testing.T) {
	_, err := NewProduct("id", "catid", "", "desc", 10.0, true)
	require.Error(t, err)
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	_, err := NewProduct("id", "catid", "Coca-Cola", "desc", -1.0, true)
	require.Error(t, err)
}

func TestProduct_SetPrice_Valid(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	err := p.SetPrice(10.0)
	require.NoError(t, err)
	require.Equal(t, 10.0, p.Price.Value())
}

func TestProduct_SetPrice_Invalid(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	err := p.SetPrice(-1.0)
	require.Error(t, err)
}

func TestProduct_SetDescription(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	err := p.SetDescription("nova desc")
	require.NoError(t, err)
}

func TestProduct_Activate_Deactivate(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, false)
	err := p.Activate()
	require.NoError(t, err)
	require.True(t, p.Active)
	err = p.Deactivate()
	require.NoError(t, err)
	require.False(t, p.Active)
}

func TestProduct_SetCategory(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	err := p.SetCategory("nova-cat")
	require.NoError(t, err)
	require.Equal(t, "nova-cat", p.CategoryID)
}

func TestProduct_AddImage_Valid(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	fileName, err := p.AddImage("img.jpg")
	require.NoError(t, err)
	require.NotNil(t, fileName)
}

func TestProduct_AddImage_Invalid(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	_, err := p.AddImage("")
	require.Error(t, err)
}

func TestProduct_RemoveImage_LastImage(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	err := p.RemoveImage(p.Images[0].FileName)
	require.Error(t, err)
}

func TestProduct_RemoveImage_NotFound(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	_, _ = p.AddImage("img2.jpg")
	err := p.RemoveImage("naoexiste.jpg")
	require.Error(t, err)
}

func TestProduct_RemoveImage_Success(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	fileName, _ := p.AddImage("img2.jpg")
	err := p.RemoveImage(*fileName)
	require.NoError(t, err)
}

func TestProduct_ImageIsDefault(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	isDefault := p.ImageIsDefault(p.Images[0].FileName)
	require.True(t, isDefault)
	_, _ = p.AddImage("img2.jpg")
	isDefault = p.ImageIsDefault("img2.jpg")
	require.False(t, isDefault)
}

func TestProduct_IsEmpty(t *testing.T) {
	p := &Product{}
	require.True(t, p.IsEmpty())
	p.ID = "id"
	require.False(t, p.IsEmpty())
}

func TestProduct_SetAllPreviousImagesAsNotDefault(t *testing.T) {
	p, _ := NewProduct("id", "catid", "Coca-Cola", "desc", 5.99, true)
	_, _ = p.AddImage("img2.jpg")
	p.SetAllPreviousImagesAsNotDefault()
	require.False(t, p.Images[0].IsDefault)
}

func TestNewProductWithImages_Valid(t *testing.T) {
	images := []struct{ FileName, Url string }{
		{"img1.jpg", "http://localhost/img1.jpg"},
		{"img2.jpg", "http://localhost/img2.jpg"},
	}
	p, err := NewProductWithImages("id", "catid", "Coca-Cola", "desc", 5.99, true, images)
	require.NoError(t, err)
	require.Equal(t, "id", p.ID)
	require.Equal(t, "catid", p.CategoryID)
	require.Equal(t, "Coca-Cola", p.Name.Value())
	require.Equal(t, "desc", p.Description)
	require.Equal(t, 5.99, p.Price.Value())
	require.Equal(t, 2, len(p.Images))
	require.True(t, p.Active)
}

func TestNewProductWithImages_InvalidName(t *testing.T) {
	images := []struct{ FileName, Url string }{
		{"img1.jpg", "http://localhost/img1.jpg"},
	}
	_, err := NewProductWithImages("id", "catid", "", "desc", 5.99, true, images)
	require.Error(t, err)
}

func TestNewProductWithImages_InvalidPrice(t *testing.T) {
	images := []struct{ FileName, Url string }{
		{"img1.jpg", "http://localhost/img1.jpg"},
	}
	_, err := NewProductWithImages("id", "catid", "Coca-Cola", "desc", -1.0, true, images)
	require.Error(t, err)
}

func TestNewProductWithImages_InvalidImage(t *testing.T) {
	images := []struct{ FileName, Url string }{
		{"", "http://localhost/img1.jpg"}, // FileName inválido
	}
	_, err := NewProductWithImages("id", "catid", "Coca-Cola", "desc", 5.99, true, images)
	require.Error(t, err)
}

func TestNewProductImage_And_ToDAO(t *testing.T) {
	img, _ := value_objects.NewImage("img1.jpg")
	createdAt := time.Now()
	prodImg := NewProductImage("imgid", "prodid", img, createdAt, true)

	require.Equal(t, "imgid", prodImg.ID)
	require.Equal(t, "prodid", prodImg.ProductID)
	require.Equal(t, img.FileName, prodImg.FileName)
	require.Equal(t, img.Url, prodImg.Url)
	require.Equal(t, createdAt, prodImg.CreatedAt)
	require.True(t, prodImg.IsDefault)

	dao := prodImg.ToDAO()
	require.Equal(t, prodImg.ID, dao.ID)
	require.Equal(t, prodImg.ProductID, dao.ProductID)
	require.Equal(t, prodImg.FileName, dao.FileName)
	require.Equal(t, prodImg.Url, dao.Url)
	require.Equal(t, prodImg.CreatedAt, dao.CreatedAt)
	require.Equal(t, prodImg.IsDefault, dao.IsDefault)
}
