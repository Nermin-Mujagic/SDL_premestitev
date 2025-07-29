package priorityqueue

import (
	"testing"

	"premestitev.sdl/v2/internal/helpers"
	request "premestitev.sdl/v2/internal/requests"
)

var (
	genericRequest, _ = request.CreateTransferRequest("nmujag", nil, false, nil, false, nil)
	emptyPriorityList = PriorityList{}
)

func TestAddRequest(t *testing.T) {
	t.Run("valid add to prio", func(t *testing.T) {
		prio := emptyPriorityList
		err := prio.AddRequest(*genericRequest)
		helpers.AssertNoError(t, err)
		expected := PriorityList{*genericRequest}
		helpers.AssertDeepEqual(t, prio, expected)
	})

	t.Run("duplicate add to prio", func(t *testing.T) {
		prio := emptyPriorityList
		err := prio.AddRequest(*genericRequest)
		helpers.AssertNoError(t, err)

		err = prio.AddRequest(*genericRequest)
		helpers.AssertError(t, err)

		expected := PriorityList{*genericRequest}
		helpers.AssertDeepEqual(t, prio, expected)

	})

}
