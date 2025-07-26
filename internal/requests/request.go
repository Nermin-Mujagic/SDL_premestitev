package request

import (
	"errors"
	"fmt"
	"math/rand"
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

type TransferRequest struct {
	RequestID      int
	StudentID      string
	PreferredDorms []string // "FDV, Poljane, I,II,III..."
	Apartment      bool
	RoomType       *string // "any, full, single, couple"
	WithPartner    bool
	PartnerID      *string
	DateSubmitted  time.Time
	Status         RequestStatus // "pending (0)", "approved (1)", "declined (2)"
}

func (tr TransferRequest) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "\nRequest{ID:%d, Student:%s", tr.RequestID, tr.StudentID)

	if tr.WithPartner && tr.PartnerID != nil {
		fmt.Fprintf(&sb, ", Partner:%s", *tr.PartnerID)
	}

	switch len(tr.PreferredDorms) {
	case 0:
		fmt.Fprint(&sb, "\n Dorms: any")
	case 1:
		fmt.Fprintf(&sb, "\n Dorms: %s", tr.PreferredDorms[0])
	default:
		fmt.Fprintf(&sb, "\n Dorms: %v", tr.PreferredDorms)
	}

	if tr.RoomType != nil {
		fmt.Fprintf(&sb, ", RoomType: %s", *tr.RoomType)
	}

	fmt.Fprintf(&sb, "}")

	return sb.String()

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

	r := rand.New(rand.NewSource(47))

	newRequest := TransferRequest{
		RequestID:      r.Intn(1000),
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
