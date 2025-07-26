package priorityqueue

import (
	"fmt"

	request "premestitev.sdl/v2/internal/requests"
)

type PriorityList []request.TransferRequest

func (pl *PriorityList) AddRequest(req request.TransferRequest) error {
	for _, existing := range *pl {
		if existing.StudentID == req.StudentID {
			return fmt.Errorf("student %s already has active request", req.StudentID)
		}
	}
	*pl = append(*pl, req)
	return nil
}
