package main

import (
	"log"

	request "premestitev.sdl/v2/requests"
)

var ()

func main() {
	var (
		cpl = "coupleApartment"
		// single = "single"
	)

	_, err := request.CreateTransferRequest("", []string{}, false, nil, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a very simple transfer request
	r1, err := request.CreateTransferRequest("nmujag", []string{"any"}, false, nil, false, nil)

	if err != nil {
		log.Fatal(err)
	}

	partner2 := "tpriimek"
	r2, err := request.CreateTransferRequest("jdorn", []string{"FDV", "Poljane"}, true, &cpl, true, &partner2)

	if err != nil {
		log.Fatal(err)
	}

	r3, err := request.CreateTransferRequest("twajs", []string{"Mestni Log", "Poljane", "Gerbiceva"}, false, nil, false, nil)

	if err != nil {
		log.Fatal(err)
	}
	allRequests := request.TransferRequests{*r1, *r2, *r3}
	request.PrintRequests(allRequests)

}
