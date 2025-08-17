package aggregates

import (
	"time"

	"github.com/google/uuid"
)

// Patient represents a patient in the FHIR system.
// PURPOSE: Aggregate root that encapsulates all business rules related to the patient.
// HIPAA COMPLIANCE: Contains PHI (Protected Health Information) that must be encrypted.
type Patient struct {
	// Unique identification
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FHIRId   string    `json:"fhir_id" gorm:"unique;not null"`            // ID do FHIR server
	TenantID uuid.UUID `json:"tenant_id" gorm:"type:uuid;not null;index"` // Multi-tenancy

	// === PHI - SENSITIVE DATA (encrypted at rest) ===
	// Personal identification
	FirstName string     `json:"first_name,omitempty" gorm:"encrypted"`
	LastName  string     `json:"last_name,omitempty" gorm:"encrypted"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
	Gender    string     `json:"gender,omitempty"` // male, female, other, unknown

	// Contact
	PhoneNumbers []string `json:"phone_numbers,omitempty" gorm:"type:text[];encrypted"`
	EmailAddress string   `json:"email_address,omitempty" gorm:"encrypted"`

	// Address
	Addresses []Address `json:"addresses,omitempty" gorm:"type:jsonb;encrypted"`

	// === NON-SENSITIVE METADATA ===
	// Status and control
	Active    bool      `json:"active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Audit and compliance
	Version   int    `json:"version" gorm:"default:1"` // Optimistic locking
	CreatedBy string `json:"created_by"`               // User who created
	UpdatedBy string `json:"updated_by"`               // Last user who updated

	// Relationships (not loaded by default)
	Encounters []Encounter `json:"encounters,omitempty" gorm:"foreignKey:PatientID"`
}

// Address represents a patient's address
// PURPOSE: Value object that is part of the Patient aggregate
type Address struct {
	Use        string `json:"use,omitempty"`  // home, work, temp, old
	Type       string `json:"type,omitempty"` // postal, physical, both
	Line       string `json:"line,omitempty"` // Street, number, additional information
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`
	Country    string `json:"country,omitempty"`
}

// === BUSINESS METHODS ===

// NewPatient creates a new patient by validating business rules.
// RULE: Every patient must have at least a FirstName and TenantID.
func NewPatient(tenantID uuid.UUID, fhirID, firstName string) (*Patient, error) {
	if tenantID == uuid.Nil {
		return nil, ErrInvalidTenantID
	}
	if fhirID == "" {
		return nil, ErrInvalidFHIRID
	}
	if firstName == "" {
		return nil, ErrInvalidFirstName
	}

	return &Patient{
		ID:        uuid.New(),
		TenantID:  tenantID,
		FHIRId:    fhirID,
		FirstName: firstName,
		Active:    true,
		Version:   1,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
