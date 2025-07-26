package priorityqueue

import request "premestitev.sdl/v2/internal/requests"

type PriorityList []request.TransferRequest

func (p *PriorityList) AddRequest(r request.TransferRequest) {
	*p = append(*p, r)
}
