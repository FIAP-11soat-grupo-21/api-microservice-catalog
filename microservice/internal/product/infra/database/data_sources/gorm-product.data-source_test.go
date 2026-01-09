package data_sources_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/data_sources"
)

func TestGormProductDataSource_Insert(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "products"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Insert(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true})
	require.NoError(t, err)
}

func TestGormProductDataSource_Insert_WithImages(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	// Expectação para inserir o produto
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "products"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Expectação para inserir a imagem do produto
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "product_images"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	img := daos.ProductImageDAO{ID: "imgid1", FileName: "img.jpg"}
	product := daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true, Images: []daos.ProductImageDAO{img}}
	err := ds.Insert(product)
	require.NoError(t, err)
}

func TestGormProductDataSource_Insert_ErrorOnCreate(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "products"`)).WillReturnError(errors.New("erro ao inserir produto"))
	mock.ExpectRollback()
	err := ds.Insert(daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true})
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao inserir produto")
}

func TestGormProductDataSource_Insert_ErrorOnAddProductImage(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	// Cria um data source com um mock do método AddProductImage que retorna erro
	ds := data_sources.NewProductDataSource(db)
	// Espera a inserção do produto
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "products"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	// Espera a inserção da imagem e retorna erro
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "product_images"`)).WillReturnError(errors.New("erro ao inserir imagem"))
	mock.ExpectRollback()
	img := daos.ProductImageDAO{ID: "imgid1", FileName: "img.jpg"}
	product := daos.ProductDAO{ID: "pid", Name: "Produto Teste", Description: "desc", Price: 10.0, CategoryID: "cat1", Active: true, Images: []daos.ProductImageDAO{img}}
	err := ds.Insert(product)
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao inserir imagem")
}

func TestGormProductDataSource_FindAll(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category_id", "active"}).AddRow("pid", "Produto Teste", "desc", 10.0, "cat1", true)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products"`)).WillReturnRows(rows)
	// Expectação para busca de imagens do produto
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_images" WHERE "product_images"."product_id" = $1 AND is_default = $2 ORDER BY created_at desc`)).WithArgs("pid", true).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "file_name", "url", "is_default", "created_at"}))
	products, err := ds.FindAll()
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "pid", products[0].ID)
}

func TestGormProductDataSource_FindAll_ErrorOnQuery(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products"`)).WillReturnError(errors.New("erro ao buscar produtos"))
	products, err := ds.FindAll()
	require.Error(t, err)
	require.Nil(t, products)
	require.Contains(t, err.Error(), "erro ao buscar produtos")
}

func TestGormProductDataSource_FindAllByCategoryID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category_id", "active"}).AddRow("pid", "Produto Teste", "desc", 10.0, "cat1", true)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE category_id = $1`)).WithArgs("cat1").WillReturnRows(rows)
	// Expectação para busca de imagens do produto
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_images" WHERE "product_images"."product_id" = $1 AND is_default = $2 ORDER BY created_at desc`)).WithArgs("pid", true).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "file_name", "url", "is_default", "created_at"}))
	products, err := ds.FindAllByCategoryID("cat1")
	require.NoError(t, err)
	require.Len(t, products, 1)
	require.Equal(t, "cat1", products[0].CategoryID)
}

func TestGormProductDataSource_FindAllByCategoryID_ErrorOnQuery(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE category_id = $1`)).WithArgs("cat1").WillReturnError(errors.New("erro ao buscar produtos por categoria"))
	products, err := ds.FindAllByCategoryID("cat1")
	require.Error(t, err)
	require.Nil(t, products)
	require.Contains(t, err.Error(), "erro ao buscar produtos por categoria")
}

func TestGormProductDataSource_FindByID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "category_id", "active"}).AddRow("pid", "Produto Teste", "desc", 10.0, "cat1", true)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE id = $1 ORDER BY "products"."id" LIMIT $2`)).WithArgs("pid", 1).WillReturnRows(rows)
	// Expectação para busca das imagens do produto
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_images" WHERE "product_images"."product_id" = $1 AND is_default = $2`)).WithArgs("pid", true).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "file_name", "url", "is_default", "created_at"}))
	product, err := ds.FindByID("pid")
	require.NoError(t, err)
	require.Equal(t, "pid", product.ID)
}

func TestGormProductDataSource_FindByID_NotFound(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "products" WHERE id = $1 ORDER BY "products"."id" LIMIT $2`)).WithArgs("pid404", 1).WillReturnError(gorm.ErrRecordNotFound)
	_, err := ds.FindByID("pid404")
	require.Error(t, err)
}

func TestGormProductDataSource_Update(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "products"`)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Update(daos.ProductDAO{ID: "pid", Name: "Produto Atualizado", Description: "desc", Price: 20.0, CategoryID: "cat1", Active: true})
	require.NoError(t, err)
}

func TestGormProductDataSource_Update_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "products"`)).WillReturnError(errors.New("erro ao atualizar produto"))
	mock.ExpectRollback()
	err := ds.Update(daos.ProductDAO{ID: "pid", Name: "Produto Atualizado", Description: "desc", Price: 20.0, CategoryID: "cat1", Active: true})
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao atualizar produto")
}

func TestGormProductDataSource_Delete(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "products" WHERE id = $1`)).WithArgs("pid").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Delete("pid")
	require.NoError(t, err)
}

func TestGormProductDataSource_Delete_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "products" WHERE id = $1`)).WithArgs("pid").WillReturnError(errors.New("erro ao deletar produto"))
	mock.ExpectRollback()
	err := ds.Delete("pid")
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao deletar produto")
}

func TestGormProductDataSource_SetAllPreviousImagesAsNotDefault(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2 AND id <> $3`)).WithArgs(false, "pid", "imgid2").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.SetAllPreviousImagesAsNotDefault("pid", "imgid2")
	require.NoError(t, err)
}

func TestGormProductDataSource_SetAllPreviousImagesAsNotDefault_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2 AND id <> $3`)).WithArgs(false, "pid", "imgid2").WillReturnError(errors.New("erro ao atualizar imagens"))
	mock.ExpectRollback()
	err := ds.SetAllPreviousImagesAsNotDefault("pid", "imgid2")
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao atualizar imagens")
}

func TestGormProductDataSource_FindAllImagesProductById(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	// Use time.Time para o campo created_at
	timeNow := time.Now()
	rows := sqlmock.NewRows([]string{"id", "product_id", "file_name", "url", "is_default", "created_at"}).AddRow("imgid1", "pid", "img.jpg", "url", true, timeNow)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_images" WHERE product_id = $1 ORDER BY created_at desc`)).WithArgs("pid").WillReturnRows(rows)
	images, err := ds.FindAllImagesProductById("pid")
	require.NoError(t, err)
	require.Len(t, images, 1)
	require.Equal(t, "imgid1", images[0].ID)
}

func TestGormProductDataSource_FindAllImagesProductById_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "product_images" WHERE product_id = $1 ORDER BY created_at desc`)).WithArgs("pid").WillReturnError(errors.New("erro ao buscar imagens"))
	images, err := ds.FindAllImagesProductById("pid")
	require.Error(t, err)
	require.Nil(t, images)
	require.Contains(t, err.Error(), "erro ao buscar imagens")
}

func TestGormProductDataSource_DeleteImage(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "product_images" WHERE file_name = $1`)).WithArgs("img.jpg").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.DeleteImage("img.jpg")
	require.NoError(t, err)
}

func TestGormProductDataSource_DeleteImage_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewProductDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "product_images" WHERE file_name = $1`)).WithArgs("img.jpg").WillReturnError(errors.New("erro ao deletar imagem"))
	mock.ExpectRollback()
	err := ds.DeleteImage("img.jpg")
	require.Error(t, err)
	require.Contains(t, err.Error(), "erro ao deletar imagem")
}

// func TestGormProductDataSource_SetImageAsDefault(t *testing.T) {
// 	db, mock, cleanup := setupMockDB(t)
// 	defer cleanup()
// 	ds := data_sources.NewProductDataSource(db)
// 	// Não espere transação, apenas os dois updates
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2`)).WithArgs(false, "pid").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2 AND id = $3`)).WithArgs(true, "pid", "imgid1").WillReturnResult(sqlmock.NewResult(1, 1))
// 	err := ds.SetImageAsDefault("pid", "imgid1")
// 	require.NoError(t, err)
// }

// func TestGormProductDataSource_SetImageAsDefault_ErrorOnUnsetAll(t *testing.T) {
// 	db, mock, cleanup := setupMockDB(t)
// 	defer cleanup()
// 	ds := data_sources.NewProductDataSource(db)
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2`)).WithArgs(false, "pid").WillReturnError(errors.New("erro ao unsetar imagens"))
// 	err := ds.SetImageAsDefault("pid", "imgid1")
// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "erro ao unsetar imagens")
// }

// func TestGormProductDataSource_SetImageAsDefault_ErrorOnSetDefault(t *testing.T) {
// 	db, mock, cleanup := setupMockDB(t)
// 	defer cleanup()
// 	ds := data_sources.NewProductDataSource(db)
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2`)).WithArgs(false, "pid").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "product_images" SET "is_default"=$1 WHERE product_id = $2 AND id = $3`)).WithArgs(true, "pid", "imgid1").WillReturnError(errors.New("erro ao setar imagem default"))
// 	err := ds.SetImageAsDefault("pid", "imgid1")
// 	require.Error(t, err)
// 	require.Contains(t, err.Error(), "erro ao setar imagem default")
// }
