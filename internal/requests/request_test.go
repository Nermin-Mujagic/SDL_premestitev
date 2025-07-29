package request

import (
	"testing"

	"premestitev.sdl/v2/internal/helpers"
)

var (
	std1 = "ipriimek"
	std2 = "jnovak"
)

func TestCreateRequest(t *testing.T) {
	t.Run("empty student id", func(t *testing.T) {
		_, err := CreateTransferRequest("", []string{}, false, nil, false, nil)
		helpers.AssertError(t, err)
	})
	t.Run("valid dorm", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{"dom_1"}, false, nil, false, nil)
		helpers.AssertNoError(t, err)
		helpers.AssertDeepEqual(t, req.PreferredDorms, []string{"dom_1"})
	})

	t.Run("invalid dorm", func(t *testing.T) {
		_, err := CreateTransferRequest(std1, []string{"dom_schmit"}, false, nil, false, nil)

		helpers.AssertError(t, err)
	})

	t.Run("empty dorm list", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, false, nil, false, nil)
		helpers.AssertNoError(t, err)
		if len(req.PreferredDorms) > 1 {
			t.Errorf("expected any in preferred dorms, %+v", req.PreferredDorms)
		}
		helpers.AssertEqualStrings(t, req.PreferredDorms[0], "any")

	})

	t.Run("no room type", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, false, nil, false, nil)
		helpers.AssertNoError(t, err)
		helpers.AssertEqual(t, req.RoomType, nil)

		emptyRoom := ""
		req, err = CreateTransferRequest(std1, []string{}, false, &emptyRoom, false, nil)
		helpers.AssertNoError(t, err)
		helpers.AssertEqual(t, req.RoomType, nil)
	})

	t.Run("invalid room type", func(t *testing.T) {
		invalidRoom := "penthouse"
		_, err := CreateTransferRequest(std1, []string{}, false, &invalidRoom, false, nil)
		helpers.AssertError(t, err)
	})

	t.Run("valid room type", func(t *testing.T) {
		validRoom := "coupleApartment"
		req, err := CreateTransferRequest(std1, []string{}, false, &validRoom, false, nil)
		helpers.AssertNoError(t, err)
		helpers.AssertEqualStrings(t, *req.RoomType, validRoom)
	})

	t.Run("partner missing", func(t *testing.T) {
		_, err := CreateTransferRequest(std1, []string{}, true, nil, true, nil)
		helpers.AssertError(t, err)
	})

	t.Run("partner valid", func(t *testing.T) {
		req, err := CreateTransferRequest(std1, []string{}, true, nil, true, &std2)
		helpers.AssertNoError(t, err)
		helpers.AssertEqualStrings(t, *req.PartnerID, "jnovak")
	})

}
