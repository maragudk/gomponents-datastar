// Package assert provides testing helpers.
package assert

import (
	"strings"
	"testing"

	. "maragu.dev/gomponents"
)

// Equal checks for equality between the given expected string and the rendered Node string.
func Equal(t *testing.T, expected string, actual Node) {
	t.Helper()

	var b strings.Builder
	err := actual.Render(&b)
	if err != nil {
		t.Fatal("error rendering actual:", err)
	}
	if expected != b.String() {
		t.Fatalf(`expected "%v" but got "%v"`, expected, b.String())
	}
}

// EqualString checks for equality between the given expected and actual strings.
func EqualString(t *testing.T, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Fatalf(`expected "%v" but got "%v"`, expected, actual)
	}
}

// Error checks for a non-nil error.
func Error(t *testing.T, err error) {
	t.Helper()

	if err == nil {
		t.Fatal("error is nil")
	}
}
