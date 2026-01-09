package suites

import (
	"context"
	"tech_challenge/bdd/steps"
	mock_interfaces "tech_challenge/internal/product/interfaces/mocks"

	"github.com/cucumber/godog"
	"github.com/golang/mock/gomock"
)

type godogReporter struct{}

func (r *godogReporter) Errorf(_ string, _ ...interface{}) {
}
func (r *godogReporter) Fatalf(format string, _ ...interface{}) { panic("gomock fatal: " + format) }

func InitializeScenario(ctx *godog.ScenarioContext) {
	var helper *steps.CategoryHelper

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		ctrl := gomock.NewController(&godogReporter{})
		mockDS := mock_interfaces.NewMockICategoryDataSource(ctrl)

		helper = &steps.CategoryHelper{
			Ctrl:   ctrl,
			MockDS: mockDS,
		}
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		if helper != nil && helper.Ctrl != nil {
			helper.Ctrl.Finish()
		}
		return ctx, nil
	})

	ctx.Step(`^the category data is valid with name "([^"]*)"$`, func(name string) error {
		helper.SetCategoryName(name)
		return helper.TheCategoryDataIsValid()
	})
	ctx.Step(`^I send a request to create a new category$`, func() error { return helper.ISendARequestToCreateANewCategory() })
	ctx.Step(`^the category should be created successfully$`, func() error { return helper.CategoryShouldBeCreated() })
}
