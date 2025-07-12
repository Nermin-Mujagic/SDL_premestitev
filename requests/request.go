package request

import (
	"fmt"
	"time"
)

type statusRequest int

const (
	Pending statusRequest = iota
	Approved
	Declined
)

type TransferRequests []TransferRequest

type TransferRequest struct {
	StudentID     string
	PreferredDorm []string // "FDV, Poljane, I,II,III..."
	RoomType      string   // "any, full, single, couple"
	WithPartner   bool
	PartnerID     *string
	DateSubmitted time.Time
	Status        statusRequest // "pending (0)", "approved (1)", "declined (2)"
}

func CreateTransferRequest(studentID string, preferredDorm []string, roomType string, partnerID *string) (*TransferRequest, error) {
	withPartner := false
	if partnerID != nil {
		withPartner = true
	}

	if studentID == "" {
		return nil, fmt.Errorf("Student ID is invalid")
	}

	newRequest := TransferRequest{
		StudentID:     studentID,
		PreferredDorm: preferredDorm,
		RoomType:      roomType,
		WithPartner:   withPartner,
		PartnerID:     partnerID,
		DateSubmitted: time.Now(),
		Status:        Pending,
	}

	return &newRequest, nil
}

func PrintRequests(transferRequests TransferRequests) {
	fmt.Println("--- Pending prošnje ---")

	for i, request := range transferRequests {
		if request.Status == Pending {
			fmt.Printf("Št. prošnje: %d\n", i+1)
			fmt.Printf("Vlagatelj: %s\n", request.StudentID)
			fmt.Printf("Izbrani dom(ovi): %v\n", request.PreferredDorm)
			fmt.Printf("vrsta sobe: %s\n", request.RoomType)
			if request.WithPartner {
				fmt.Printf("Z osebo: %s\n", *request.PartnerID)
			}
			fmt.Println("-----------------")
		}

	}
}
