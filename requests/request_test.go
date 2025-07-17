package request

import (
	"reflect"
	"testing"
)

const std1 = "ipriimek"

func TestCreateRequest(t *testing.T) {
	t.Run("test empty student id", func(t *testing.T) {
		_, err := CreateTransferRequest("", []string{}, false, nil, nil)
		AssertError(t, err)
	})
	t.Run("test valid dorm", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{"dom_1"}, false, nil, nil)
		AssertNoError(t, err)

		AssertDeepEqual(t, req.PreferredDorms, []string{"dom_1"})
	})

	t.Run("test invalid dorm", func(t *testing.T) {
		_, err := CreateTransferRequest(std1, []string{"dom_schmit"}, false, nil, nil)
		AssertError(t, err)
	})

}

func AssertDeepEqual(t testing.TB, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %+v, got %+v", got, want)
	}
}

func AssertEqualStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func AssertError(t testing.TB, got error) {
	if got == nil {
		t.Errorf("expected error")
	}
}

func AssertNoError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("expected no error, got %v", got)
	}
}
