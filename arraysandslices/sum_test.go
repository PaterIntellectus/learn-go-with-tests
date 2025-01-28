package arraysandslices

import (
	"fmt"
	"testing"

	"example.com/gowithtests/tester"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		tester.AssertEqualWithDetails(t, got, want, fmt.Sprintf("given %v", numbers))
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum many colection", func(t *testing.T) {
		first := []int{1, 2}
		second := []int{0, 9}

		got := SumAll(first, second)
		want := []int{3, 9}

		tester.AssertSlicesEqual(t, got, want)
	})
}

func TestSumTails(t *testing.T) {
	t.Run("sum tails", func(t *testing.T) {
		first := []int{1, 2}
		second := []int{0, 9}

		got := SumAllTails(first, second)
		want := []int{2, 9}

		tester.AssertSlicesEqual(t, got, want)
	})

	t.Run("sum tails of empty slice", func(t *testing.T) {
		first := []int{}
		second := []int{3, 4, 5}

		got := SumAllTails(first, second)
		want := []int{0, 9}

		tester.AssertSlicesEqual(t, got, want)
	})
}
