package request

import (
	"errors"
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

func CreateTransferRequest(studentID string, preferredDormList []string, apartment bool, roomType *string, withPartner bool, partnerID *string) (*TransferRequest, error) {
	if studentID == "" {
		return nil, fmt.Errorf("Student ID %q is invalid", studentID)
	}

	dormList, err := validateDormList(preferredDormList)
	if err != nil {
		return nil, err
	}

	switch {
	case roomType == nil:
	case *roomType == "":
		roomType = nil
	default:
		if !slices.Contains(RoomTypes, *roomType) {
			return nil, fmt.Errorf("invalid room type: %q", *roomType)
		}

	}

	if withPartner {
		if partnerID == nil || *partnerID == "" {
			return nil, errors.New("partner missing")
		}
	}

	newRequest := TransferRequest{
		StudentID:      studentID,
		PreferredDorms: dormList,
		Apartment:      apartment,
		RoomType:       roomType,
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
