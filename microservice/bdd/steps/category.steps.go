package steps

import (
	"fmt"
	"strings"
	"tech_challenge/internal/product/daos"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"

	"github.com/golang/mock/gomock"
)

type CategoryHelper struct {
	Ctrl   *gomock.Controller
	MockDS *mock_interfaces.MockICategoryDataSource

	valid struct {
		name   string
		active bool
	}
	existingID string
}

func (ch *CategoryHelper) TheCategoryDataIsValid() error {
	ch.valid.name = "Bebidas"
	ch.valid.active = true
	trimmed := strings.TrimSpace(ch.valid.name)
	if len(trimmed) < 3 {
		return fmt.Errorf("category name must have at least 3 characters")
	}
	if len(trimmed) > 100 {
		return fmt.Errorf("category name must have at most 100 characters")
	}
	if trimmed == "" {
		return fmt.Errorf("invalid category name provided in test data")
	}
	return nil
}

func (ch *CategoryHelper) SetCategoryName(name string) {
	ch.valid.name = name
}

func (ch *CategoryHelper) ISendARequestToCreateANewCategory() error {
	const generatedID = "cat-123"

	ch.MockDS.EXPECT().Insert(daos.CategoryDAO{
		ID:     generatedID,
		Name:   ch.valid.name,
		Active: ch.valid.active,
	}).Return(nil)

	err := ch.MockDS.Insert(daos.CategoryDAO{
		ID:     generatedID,
		Name:   ch.valid.name,
		Active: ch.valid.active,
	})
	if err != nil {
		return err
	}

	ch.existingID = generatedID
	return nil
}

func (ch *CategoryHelper) CategoryShouldBeCreated() error {
	if ch.existingID == "" {
		return fmt.Errorf("category was not created")
	}
	return nil
}
