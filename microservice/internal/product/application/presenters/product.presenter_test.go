package presenters

import (
	"os"
	"tech_challenge/internal/product/domain/entities"
	value_objects "tech_challenge/internal/product/domain/value-objects"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
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

func TestProductFromDomainToResultDTO(t *testing.T) {
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	img, _ := value_objects.NewImage("img1.jpg")
	prod := entities.Product{
		ID:          "pid",
		CategoryID:  "catid",
		Name:        name,
		Description: "desc",
		Price:       price,
		Images:      []*value_objects.Image{&img},
		Active:      true,
	}
	dto := ProductFromDomainToResultDTO(prod)
	require.Equal(t, "pid", dto.ID)
	require.Equal(t, "Coca-Cola", dto.Name)
	require.Equal(t, "desc", dto.Description)
	require.Equal(t, 5.99, dto.Price)
	require.True(t, dto.Active)
	require.Equal(t, "catid", dto.CategoryID)
	require.Len(t, dto.Images, 1)
}

func TestListProductDomainToResultDTO(t *testing.T) {
	name, _ := value_objects.NewName("Coca-Cola")
	price, _ := value_objects.NewPrice(5.99)
	img, _ := value_objects.NewImage("img1.jpg")
	prod := entities.Product{
		ID:          "pid",
		CategoryID:  "catid",
		Name:        name,
		Description: "desc",
		Price:       price,
		Images:      []*value_objects.Image{&img},
		Active:      true,
	}
	list := []entities.Product{prod}
	dtos := ListProductDomainToResultDTO(list)
	require.Len(t, dtos, 1)
	require.Equal(t, "pid", dtos[0].ID)
}

func TestProductImagesFromDomainToResultDTO(t *testing.T) {
	img, _ := value_objects.NewImage("img1.jpg")
	img2, _ := value_objects.NewImage("img2.jpg")
	imgs := []*value_objects.Image{&img, &img2}
	dtos := ProductImagesFromDomainToResultDTO(imgs)
	require.Len(t, dtos, 2)
	require.True(t, len(dtos[0].FileName) > 0 && dtos[0].FileName[:4] == "img1" && dtos[0].FileName[len(dtos[0].FileName)-4:] == ".jpg")
	require.True(t, len(dtos[1].FileName) > 0 && dtos[1].FileName[:4] == "img2" && dtos[1].FileName[len(dtos[1].FileName)-4:] == ".jpg")
}

func TestProductImageFromDomainToDTO(t *testing.T) {
	img, _ := value_objects.NewImage("img1.jpg")
	dto := ProductImageFromDomainToDTO(img)
	require.Equal(t, img.ID, dto.ID)
	require.Equal(t, img.FileName, dto.FileName)
	require.Equal(t, img.Url, dto.Url)
	require.Equal(t, img.IsDefault, dto.IsDefault)
}
