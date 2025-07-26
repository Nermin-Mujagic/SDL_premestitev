package helpers

import (
	"reflect"
	"testing"
)

func AssertEqual[K comparable](t testing.TB, got, want K) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertDeepEqual[K any](t testing.TB, got, want []K) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v\n want %v", got, want)
	}
}

func AssertEqualStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func AssertError(t testing.TB, got error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected error")
	}
	t.Log(got)
}

func AssertNoError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("expected no error, got %v", got)
	}
}
