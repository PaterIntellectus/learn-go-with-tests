package integers

import (
	"fmt"
	"testing"

	"example.com/gowithtests/tester"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	tester.AssertEqual(t, expected, sum)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
