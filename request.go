package main

import (
	"fmt"
	"time"
)

type TransferRequests []TransferRequest

type TransferRequest struct {
	StudentID     string
	PreferredDorm []string // "FDV, Poljane, I,II,III..."
	RoomType      string   // "any, full, single, couple"
	WithPartner   bool
	PartnerID     *string
	DateSubmitted time.Time
}

func CreateTransferRequest(studentID string, preferredDorm []string, roomType string, partnerID *string) TransferRequest {
	withPartner := false
	if partnerID != nil {
		withPartner = true
	}

	newRequest := TransferRequest{
		StudentID:     studentID,
		PreferredDorm: preferredDorm,
		RoomType:      roomType,
		WithPartner:   withPartner,
		PartnerID:     partnerID,
		DateSubmitted: time.Now(),
	}

	return newRequest
}

func printRequests(transferRequests TransferRequests) {
	for i, request := range transferRequests {
		fmt.Printf("Št. prošnje: %d\n", i+1)
		fmt.Printf("Vlagatelj: %s\n", request.StudentID)
		fmt.Printf("Izbrani dom(ovi): %v\n", request.PreferredDorm)
		fmt.Printf("vrsta sobe: %s\n", request.RoomType)
		if request.WithPartner {
			fmt.Printf("Z osebo: %s\n", *request.PartnerID)
		}
		fmt.Println()
	}
}

func main() {
	// Create a very simple transfer request
	r1 := CreateTransferRequest("nmujag", []string{"any"}, "any", nil)
	partner2 := "tpriimek"
	r2 := CreateTransferRequest("jdorn", []string{"FDV", "Poljane"}, "couple", &partner2)
	r3 := CreateTransferRequest("twajs", []string{"Mestni Log", "Poljane", "Gerbiceva"}, "single", nil)

	allRequests := TransferRequests{r1, r2, r3}
	printRequests(allRequests)

}
