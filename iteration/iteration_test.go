package iteration

import (
	"fmt"
	"testing"

	"example.com/gowithtests/tester"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 'a' five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		tester.AssertEqual(t, expected, repeated)
	})

	t.Run("repeat 'b' zero times", func(t *testing.T) {
		repeated := Repeat("b", 0)
		expected := ""

		tester.AssertEqual(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("c", 3)
	fmt.Println(repeated)
	// Output: ccc
}
