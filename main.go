package main

import (
	"log"

	request "premestitev.sdl/v2/requests"
)

func main() {
	// Create a very simple transfer request
	r1, err := request.CreateTransferRequest("nmujag", []string{"any"}, "any", nil)

	if err != nil {
		log.Fatal(err)
	}

	partner2 := "tpriimek"
	r2, err := request.CreateTransferRequest("jdorn", []string{"FDV", "Poljane"}, "couple", &partner2)

	if err != nil {
		log.Fatal(err)
	}

	r3, err := request.CreateTransferRequest("twajs", []string{"Mestni Log", "Poljane", "Gerbiceva"}, "single", nil)

	if err != nil {
		log.Fatal(err)
	}
	allRequests := request.TransferRequests{*r1, *r2, *r3}
	request.PrintRequests(allRequests)

}
