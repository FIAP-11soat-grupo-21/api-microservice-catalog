package controllers

import (
	"os"
	testenv "tech_challenge/internal/shared/test"
	"testing"
)

func TestMain(m *testing.M) {
	testenv.SetupTestEnv()
	code := m.Run()
	os.Exit(code)
}
