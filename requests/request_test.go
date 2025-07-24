package request

import (
	"reflect"
	"testing"
)

var (
	std1 = "ipriimek"
	std2 = "jnovak"
)

func TestCreateRequest(t *testing.T) {
	t.Run("empty student id", func(t *testing.T) {
		_, err := CreateTransferRequest("", []string{}, false, nil, false, nil)
		AssertError(t, err)
	})
	t.Run("valid dorm", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{"dom_1"}, false, nil, false, nil)
		AssertNoError(t, err)
		AssertDeepEqual(t, req.PreferredDorms, []string{"dom_1"})
	})

	t.Run("invalid dorm", func(t *testing.T) {
		_, err := CreateTransferRequest(std1, []string{"dom_schmit"}, false, nil, false, nil)

		AssertError(t, err)
	})

	t.Run("empty dorm list", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, false, nil, false, nil)
		AssertNoError(t, err)
		if len(req.PreferredDorms) > 1 {
			t.Errorf("expected any in preferred dorms, %+v", req.PreferredDorms)
		}
		AssertEqualStrings(t, req.PreferredDorms[0], "any")

	})

	t.Run("no room type", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, false, nil, false, nil)
		AssertNoError(t, err)
		AssertEqual(t, req.RoomType, nil)

		emptyRoom := ""
		req, err = CreateTransferRequest(std1, []string{}, false, &emptyRoom, false, nil)
		AssertNoError(t, err)
		AssertEqual(t, req.RoomType, nil)
	})

	t.Run("invalid room type", func(t *testing.T) {
		invalidRoom := "penthouse"
		_, err := CreateTransferRequest(std1, []string{}, false, &invalidRoom, false, nil)
		AssertError(t, err)
	})

	t.Run("valid room type", func(t *testing.T) {
		validRoom := "coupleApartment"
		req, err := CreateTransferRequest(std1, []string{}, false, &validRoom, false, nil)
		AssertNoError(t, err)
		AssertEqualStrings(t, *req.RoomType, validRoom)
	})

	t.Run("partner missing", func(t *testing.T) {
		_, err := CreateTransferRequest(std1, []string{}, true, nil, true, nil)
		AssertError(t, err)
	})

	t.Run("partner valid", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, true, nil, true, &std2)
		AssertNoError(t, err)
		AssertEqualStrings(t, *req.PartnerID, "jnovak")
	})

}

func AssertEqual[K comparable](t testing.TB, got, want K) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v, got %v", got, want)
	}
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
