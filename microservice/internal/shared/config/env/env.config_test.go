package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_IsProduction(t *testing.T) {
	c := &Config{}
	c.GoEnv = "production"
	assert.True(t, c.IsProduction())

	c.GoEnv = "development"
	assert.False(t, c.IsProduction())

	c.GoEnv = "test"
	assert.False(t, c.IsProduction())
}

func TestConfig_IsDevelopment(t *testing.T) {
	c := &Config{}
	c.GoEnv = "development"
	assert.True(t, c.IsDevelopment())

	c.GoEnv = "production"
	assert.False(t, c.IsDevelopment())

	c.GoEnv = "test"
	assert.False(t, c.IsDevelopment())
}
