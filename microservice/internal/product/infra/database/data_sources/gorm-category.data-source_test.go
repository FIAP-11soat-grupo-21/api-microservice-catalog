package data_sources_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"tech_challenge/internal/product/daos"
	"tech_challenge/internal/product/infra/database/data_sources"
)

func TestGormCategoryDataSource_Insert(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Insert(daos.CategoryDAO{ID: "cat1", Name: "Bebidas", Active: true})
	require.NoError(t, err)
}

func TestGormCategoryDataSource_FindAll(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	rows := sqlmock.NewRows([]string{"id", "name", "active"}).AddRow("cat1", "Bebidas", true)
	mock.ExpectQuery("SELECT \\* FROM \\\"category\\\"").WillReturnRows(rows)
	categories, err := ds.FindAll()
	require.NoError(t, err)
	require.Len(t, categories, 1)
	require.Equal(t, "cat1", categories[0].ID)
}

func TestGormCategoryDataSource_FindByID(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	rows := sqlmock.NewRows([]string{"id", "name", "active"}).AddRow("cat1", "Bebidas", true)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "category" WHERE id = $1 ORDER BY "category"."id" LIMIT $2`)).WithArgs("cat1", 1).WillReturnRows(rows)
	cat, err := ds.FindByID("cat1")
	require.NoError(t, err)
	require.Equal(t, "cat1", cat.ID)
}

func TestGormCategoryDataSource_FindByID_NotFound(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "category" WHERE id = $1 ORDER BY "category"."id" LIMIT $2`)).WithArgs("cat404", 1).WillReturnError(gorm.ErrRecordNotFound)
	_, err := ds.FindByID("cat404")
	require.Error(t, err)
}

func TestGormCategoryDataSource_Update(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Update(daos.CategoryDAO{ID: "cat1", Name: "Bebidas", Active: true})
	require.NoError(t, err)
}

func TestGormCategoryDataSource_Delete(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WithArgs("cat1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := ds.Delete("cat1")
	require.NoError(t, err)
}

func TestGormCategoryDataSource_Delete_Error(t *testing.T) {
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()
	ds := data_sources.NewGormCategoryDataSource(db)
	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WithArgs("cat1").WillReturnError(errors.New("delete error"))
	mock.ExpectRollback()
	err := ds.Delete("cat1")
	require.Error(t, err)
}
