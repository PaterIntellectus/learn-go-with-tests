package tester

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func AssertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	AssertEqualWithDetails(t, got, want, "")
}

func AssertEqualWithDetails(t testing.TB, got, want interface{}, details string) {
	t.Helper()

	var builder strings.Builder

	if got != want {
		builder.WriteString(fmt.Sprintf("want '%+v', but got '%+v'", want, got))

		if details != "" {
			builder.WriteString(fmt.Sprintf(", %+v", details))
		}

		t.Error(builder.String())
	}
}

func AssertDeepEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	AssertDeepEqualWithDetails(t, got, want, "")
}

func AssertDeepEqualWithDetails(t testing.TB, got, want interface{}, details string) {
	t.Helper()

	var builder strings.Builder

	if !reflect.DeepEqual(got, want) {
		builder.WriteString(fmt.Sprintf("want '%+v', but got '%+v'", want, got))

		if details != "" {
			builder.WriteString(fmt.Sprintf(", %+v", details))
		}

		t.Error(builder.String())
	}
}

func AssertSlicesEqual[S ~[]E, E comparable](t testing.TB, got, want S) {
	t.Helper()
	AssertSlicesEqualWithDetails(t, got, want, "")
}

func AssertSlicesEqualWithDetails[S ~[]E, E comparable](t testing.TB, got, want S, details string) {
	t.Helper()

	var builder strings.Builder

	if !slices.Equal(got, want) {
		builder.WriteString(fmt.Sprintf("want '%+v', but got '%+v'", want, got))

		if details != "" {
			builder.WriteString(fmt.Sprintf(", %+v", details))
		}

		t.Error(builder.String())
	}
}
