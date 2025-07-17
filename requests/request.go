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
	StudentID      string
	PreferredDorms []string // "FDV, Poljane, I,II,III..."
	RoomType       *string  // "any, full, single, couple"
	WithPartner    bool
	PartnerID      *string
	DateSubmitted  time.Time
	Status         statusRequest // "pending (0)", "approved (1)", "declined (2)"
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

func CreateTransferRequest(studentID string, preferredDormList []string, roomType *string, partnerID *string) (*TransferRequest, error) {
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

	newRequest := TransferRequest{
		StudentID:      studentID,
		PreferredDorms: dormList,
		RoomType:       roomType,
		WithPartner:    withPartner,
		PartnerID:      partnerID,
		DateSubmitted:  time.Now(),
		Status:         Pending,
	}

	return &newRequest, nil
}

func PrintRequests(transferRequests TransferRequests) {
	fmt.Println("--- Pending prošnje ---")

	for i, request := range transferRequests {
		if request.Status == Pending {
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
