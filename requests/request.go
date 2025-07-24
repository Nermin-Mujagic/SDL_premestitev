package request

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"
)

type RequestStatus string

const (
	RequestActive   RequestStatus = "active"
	RequestInactive RequestStatus = "inactive"
)

var (
	RoomTypes = []string{"singleBed", "doubleBed", "coupleApartment"}
)

type RequestValidationError struct {
	Field   string
	Message string
}

func (e RequestValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s\n%v", e.Field, e.Message)
}

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
	var invalidDorms []string

	for _, dorm := range preferredDormList {
		if _, ok := dormList[dorm]; !ok {
			invalidDorms = append(invalidDorms, dorm)
		}
	}

	if len(invalidDorms) > 0 {
		message := fmt.Sprintf("invalid dorms: %s", strings.Join(invalidDorms, ", "))
		return nil, RequestValidationError{Field: "Preferred dorms", Message: message}
	}

	return preferredDormList, nil
}

func CreateTransferRequest(studentID string, preferredDormList []string, apartment bool, roomType *string, withPartner bool, partnerID *string) (*TransferRequest, error) {
	if studentID == "" {
		return nil, &RequestValidationError{Field: "StudentID", Message: "Cannot be empty"}
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

	if withPartner && (partnerID == nil || *partnerID == "") {
		return nil, errors.New("partner missing")

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
