package request

import (
	"fmt"
	"slices"
	"time"
)

type RequestStatus string

const (
	RequestActive   RequestStatus = "active"
	RequestInactive RequestStatus = "inactive"
)

var RoomTypes = []string{"singleBed", "doubleBed", "coupleApartment"}

type TransferRequests []TransferRequest

type TransferRequest struct {
	StudentID      string
	PreferredDorms []string // "FDV, Poljane, I,II,III..."
	Apartment      bool
	RoomType       *string // "any, full, single, couple"
	WithPartner    bool
	PartnerID      *string
	DateSubmitted  time.Time
	Status         RequestStatus // "pending (0)", "approved (1)", "declined (2)"
}

func validateDormList(preferredDormList []string) ([]string, error) {
	if len(preferredDormList) == 0 {
		return []string{"any"}, nil
	}

	for _, dorm := range preferredDormList {
		if _, ok := dormList[dorm]; !ok {
			return nil, fmt.Errorf("Dorm %s is not a valid dorm", dorm)
		}
	}

	return preferredDormList, nil

}

func CreateTransferRequest(studentID string, preferredDormList []string, apartment bool, roomType *string, partnerID *string) (*TransferRequest, error) {
	withPartner := false
	if partnerID != nil {
		withPartner = true
	}

	if studentID == "" {
		return nil, fmt.Errorf("Student ID %q is invalid", studentID)
	}

	dormList, err := validateDormList(preferredDormList)
	if err != nil {
		return nil, err
	}

	var preferredRoom string
	switch roomType {
	case nil:
		preferredRoom = "any"
	default:
		if !slices.Contains(RoomTypes, *roomType) {
			return nil, fmt.Errorf("room type %q is invalid", *roomType)
		}
		preferredRoom = *roomType

	}

	newRequest := TransferRequest{
		StudentID:      studentID,
		PreferredDorms: dormList,
		Apartment:      apartment,
		RoomType:       &preferredRoom,
		WithPartner:    withPartner,
		PartnerID:      partnerID,
		DateSubmitted:  time.Now(),
		Status:         RequestActive,
	}

	return &newRequest, nil
}

func PrintRequests(transferRequests TransferRequests) {
	fmt.Println("--- Pending prošnje ---")

	for i, request := range transferRequests {
		if request.Status == RequestActive {
			fmt.Printf("Št. prošnje: %d\n", i+1)
			fmt.Printf("Vlagatelj: %s\n", request.StudentID)
			fmt.Printf("Izbrani dom(ovi): %v\n", request.PreferredDorms)
			fmt.Printf("vrsta sobe: %v\n", request.RoomType)
			if request.WithPartner {
				fmt.Printf("Z osebo: %s\n", *request.PartnerID)
			}
			fmt.Println("-----------------")
		}

	}
}
