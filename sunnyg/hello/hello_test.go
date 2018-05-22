package hello

import "testing"

func Assert(t *testing.T, actual string, expected string) {
	if (actual != expected) {
		t.Errorf("actual: '%s', expected: '%s'", actual, expected)
	}
}

func TestHello(t *testing.T) {
	t.Run("should say hello to people", func(t *testing.T) {
		actual := Hello("Chris", langEN)
		expected := "Hello, Chris"
		Assert(t, actual, expected)
	})

	t.Run("should say `Hello, World` when empty string is provided",
		func(t *testing.T) {
			actual := Hello("", langEN)
			expected := "Hello, World"
			Assert(t, actual, expected)
		})

	t.Run("should say hello in Spanish", func(t *testing.T) {
		actual := Hello("Juan", langES)
		expected := "Hola, Juan"
		Assert(t, actual, expected)
	})

	t.Run("should say hello in French", func(t *testing.T) {
		actual := Hello("Claude", langFR)
		expected := "Bonjour, Claude"
		Assert(t, actual, expected)
	})
}
