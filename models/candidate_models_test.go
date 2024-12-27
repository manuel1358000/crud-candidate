package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"crud-candidates/dtos"
)

func TestMapCandidateDtoFromCandidate(t *testing.T) {

	candidate := &Candidates{
		ID:     1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Gender: "Male",
		Salary: 50000.00,
	}

	candidateDto := candidate.MapCandidateDtoFromCandidate()

	assert.Equal(t, candidate.ID, candidateDto.ID, "El ID no coincide")
	assert.Equal(t, candidate.Name, candidateDto.Name, "El nombre no coincide")
	assert.Equal(t, candidate.Email, candidateDto.Email, "El email no coincide")
	assert.Equal(t, candidate.Gender, candidateDto.Gender, "El género no coincide")
	assert.Equal(t, candidate.Salary, candidateDto.Salary, "El salario no coincide")
}

func TestMapCandidateDtoFromCandidates(t *testing.T) {
	candidates := ListOfCandidates{
		&Candidates{
			ID:     1,
			Name:   "John Doe",
			Email:  "john.doe@example.com",
			Gender: "Male",
			Salary: 50000.00,
		},
		&Candidates{
			ID:     2,
			Name:   "Jane Smith",
			Email:  "jane.smith@example.com",
			Gender: "Female",
			Salary: 60000.00,
		},
	}

	candidateDtos := candidates.MapCandidateDtoFromCandidates()

	assert.Len(t, candidateDtos, 2, "El número de CandidateDto no coincide con el número de Candidates")

	assert.Equal(t, candidates[0].ID, candidateDtos[0].ID, "El ID del primer candidato no coincide")
	assert.Equal(t, candidates[0].Name, candidateDtos[0].Name, "El nombre del primer candidato no coincide")
	assert.Equal(t, candidates[0].Email, candidateDtos[0].Email, "El email del primer candidato no coincide")
	assert.Equal(t, candidates[0].Gender, candidateDtos[0].Gender, "El género del primer candidato no coincide")
	assert.Equal(t, candidates[0].Salary, candidateDtos[0].Salary, "El salario del primer candidato no coincide")

	assert.Equal(t, candidates[1].ID, candidateDtos[1].ID, "El ID del segundo candidato no coincide")
	assert.Equal(t, candidates[1].Name, candidateDtos[1].Name, "El nombre del segundo candidato no coincide")
	assert.Equal(t, candidates[1].Email, candidateDtos[1].Email, "El email del segundo candidato no coincide")
	assert.Equal(t, candidates[1].Gender, candidateDtos[1].Gender, "El género del segundo candidato no coincide")
	assert.Equal(t, candidates[1].Salary, candidateDtos[1].Salary, "El salario del segundo candidato no coincide")
}

func TestNewCandidate(t *testing.T) {
	candidateDto := &dtos.CandidateDto{
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Gender: "Male",
		Salary: 50000.00,
	}

	candidate := NewCandidate(candidateDto)

	assert.Equal(t, candidateDto.Name, candidate.Name, "El nombre no coincide")
	assert.Equal(t, candidateDto.Email, candidate.Email, "El email no coincide")
	assert.Equal(t, candidateDto.Gender, candidate.Gender, "El género no coincide")
	assert.Equal(t, candidateDto.Salary, candidate.Salary, "El salario no coincide")
}
