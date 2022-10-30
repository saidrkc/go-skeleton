//go:build unit
// +build unit

package app_test

import (
	"testing"

	"go-skeleton/cmd/app"
)

func TestNewEngine(t *testing.T) {
	t.Run(("Execute new engine"), func() {
		app.NewEngine()
	})
}
