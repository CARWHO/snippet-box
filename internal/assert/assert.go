package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actualValue, expectedValue T) {
	t.Helper()

	if expectedValue != actualValue {
		t.Errorf("got %v; want %v", actualValue, expectedValue)
	}
}

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()

	if !strings.Contains(actual, expectedSubstring) {
		t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}

func NilError(t *testing.T, actual error) {
	t.Helper()

	if actual != nil {
		t.Errorf("got: %v; expected: nil", actual)
	}
}
