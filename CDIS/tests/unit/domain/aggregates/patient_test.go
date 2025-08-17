package aggregates_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jasonrosa92/CDIS/CDIS/internal/domain/aggregates"
)

// TestNewPatientValidData verifica se a função NewPatient cria um paciente válido.
func TestNewPatientValidData(t *testing.T) {
	tenantID := uuid.New()
	fhirID := "patient-123"
	firstName := "John"

	patient, err := aggregates.NewPatient(tenantID, fhirID, firstName)

	if err != nil {
		t.Fatalf("esperava nenhum erro, mas obteve: %v", err)
	}
	if patient == nil {
		t.Fatal("esperava que paciente não fosse nil")
	}

	if patient.FHIRId != fhirID {
		t.Errorf("esperava FHIRId %s, mas obteve %s", fhirID, patient.FHIRId)
	}
	if patient.FirstName != firstName {
		t.Errorf("esperava FirstName %s, mas obteve %s", firstName, patient.FirstName)
	}
	if patient.TenantID != tenantID {
		t.Errorf("esperava TenantID %s, mas obteve %s", tenantID, patient.TenantID)
	}
	if patient.Active != true {
		t.Error("esperava que o paciente estivesse ativo por padrão")
	}
	if patient.Version != 1 {
		t.Errorf("esperava versão inicial 1, mas obteve %d", patient.Version)
	}
	if patient.CreatedAt.IsZero() || patient.UpdatedAt.IsZero() {
		t.Error("esperava que os timestamps de criação e atualização estivessem definidos")
	}
}

// TestNewPatientInvalidData verifica se a função NewPatient retorna erros para dados inválidos.
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
			name:        "TenantID inválido",
			tenantID:    uuid.Nil,
			fhirID:      fhirID,
			firstName:   firstName,
			expectedErr: aggregates.ErrInvalidTenantID,
		},
		{
			name:        "FHIRId vazio",
			tenantID:    tenantID,
			fhirID:      "",
			firstName:   firstName,
			expectedErr: aggregates.ErrInvalidFHIRID,
		},
		{
			name:        "FirstName vazio",
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
				t.Fatal("esperava um erro, mas obteve nil")
			}
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("esperava erro '%s', mas obteve '%s'", tc.expectedErr, err)
			}
		})
	}
}
