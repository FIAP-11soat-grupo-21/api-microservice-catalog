package database_errors_test

import (
	"errors"
	"testing"

	"tech_challenge/internal/product/domain/exceptions"
	"tech_challenge/internal/product/infra/database/database_errors"

	"github.com/stretchr/testify/require"
)

func TestHandleDatabaseErrors_NilError(t *testing.T) {
	err := database_errors.HandleDatabaseErrors(nil)
	require.NoError(t, err)
}

func TestHandleDatabaseErrors_NoSQLState(t *testing.T) {
	err := errors.New("erro generico sem sqlstate")
	result := database_errors.HandleDatabaseErrors(err)
	require.Equal(t, err, result)
}

func TestHandleDatabaseErrors_CategoryHasProducts23001(t *testing.T) {
	err := errors.New("pq: violação de restrição de integridade - SQLSTATE 23001")
	result := database_errors.HandleDatabaseErrors(err)
	_, ok := result.(*exceptions.CategoryHasProductsException)
	require.True(t, ok)
}

func TestHandleDatabaseErrors_CategoryHasProducts23503(t *testing.T) {
	err := errors.New("pq: violação de restrição de integridade - SQLSTATE 23503")
	result := database_errors.HandleDatabaseErrors(err)
	_, ok := result.(*exceptions.CategoryHasProductsException)
	require.True(t, ok)
}

func TestHandleDatabaseErrors_OtherSQLState(t *testing.T) {
	err := errors.New("erro qualquer - SQLSTATE 99999")
	result := database_errors.HandleDatabaseErrors(err)
	require.Equal(t, err, result)
}

func TestExtractDatabaseState(t *testing.T) {
	msg := "erro - SQLSTATE 23503"
	code := database_errors.ExtractDatabaseState(msg)
	require.Equal(t, "23503", code)

	msg2 := "erro sem sqlstate"
	code2 := database_errors.ExtractDatabaseState(msg2)
	require.Equal(t, "", code2)
}
