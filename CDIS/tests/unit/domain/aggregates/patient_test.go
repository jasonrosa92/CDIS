package aggregates_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jasonrosa92/CDIS/CDIS/internal/domain/aggregates"
)

// TestNewPatientValidData checks whether the NewPatient function creates a valid patient.
func TestNewPatientValidData(t *testing.T) {
	tenantID := uuid.New()
	fhirID := "patient-123"
	firstName := "John"

	patient, err := aggregates.NewPatient(tenantID, fhirID, firstName)

	if err != nil {
		t.Fatalf("I expected no errors, but got: %v", err)
	}
	if patient == nil {
		t.Fatal("I hoped that the patient wasn't nil.")
	}

	if patient.FHIRId != fhirID {
		t.Errorf("Expected FHIRId %s, but got %s", fhirID, patient.FHIRId)
	}
	if patient.FirstName != firstName {
		t.Errorf("Expcted FirstName %s, but got %s", firstName, patient.FirstName)
	}
	if patient.TenantID != tenantID {
		t.Errorf("Expected TenantID %s, but got %s", tenantID, patient.TenantID)
	}
	if patient.Active != true {
		t.Error("I expected the patient to be active by default.")
	}
	if patient.Version != 1 {
		t.Errorf("I was expecting initial version 1, but got %d", patient.Version)
	}
	if patient.CreatedAt.IsZero() || patient.UpdatedAt.IsZero() {
		t.Error("I expected the creation and update timestamps to be set.")
	}
}

// TestNewPatientInvalidData checks whether the NewPatient function returns errors for invalid data.
func TestNewPatientInvalidData(t *testing.T) {
	tenantID := uuid.New()
	fhirID := "patient-123"
	firstName := "John"

	testCases := []struct {
		name        string
		tenantID    uuid.UUID
		fhirID      string
		firstName   string
		expectedErr error
	}{
		{
			name:        "TenantID invalid",
			tenantID:    uuid.Nil,
			fhirID:      fhirID,
			firstName:   firstName,
			expectedErr: aggregates.ErrInvalidTenantID,
		},
		{
			name:        "FHIRId empty",
			tenantID:    tenantID,
			fhirID:      "",
			firstName:   firstName,
			expectedErr: aggregates.ErrInvalidFHIRID,
		},
		{
			name:        "FirstName empty",
			tenantID:    tenantID,
			fhirID:      fhirID,
			firstName:   "",
			expectedErr: aggregates.ErrInvalidFirstName,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := aggregates.NewPatient(tc.tenantID, tc.fhirID, tc.firstName)
			if err == nil {
				t.Fatal("Expected an ERROR, but got nil")
			}
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf(" Expected an ERROR '%s', but got '%s'", tc.expectedErr, err)
			}
		})
	}
}
