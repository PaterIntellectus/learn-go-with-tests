package helloworld

import (
	"testing"

	"example.com/gowithtests/tester"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Pater", russian)
		want := "Привет, Pater!"

		tester.AssertEqual(t, got, want)
	})

	t.Run("empty string default to 'World'", func(t *testing.T) {
		got := Hello("", spanish)
		want := "Hola, World!"

		tester.AssertEqual(t, got, want)
	})

	t.Run("empty of unsupported languages defaults to english", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		tester.AssertEqual(t, got, want)
	})

	t.Run("french greeeting", func(t *testing.T) {
		got := Hello("Satura", french)
		want := "Bonjour, Satura!"

		tester.AssertEqual(t, got, want)
	})
}
