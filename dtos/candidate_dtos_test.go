package dtos

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandidateDtoSerialization(t *testing.T) {
	candidate := CandidateDto{
		ID:     1,
		Name:   "John Doe",
		Email:  "john.doe@example.com",
		Gender: "Male",
		Salary: 50000.00,
	}

	jsonData, err := json.Marshal(candidate)
	assert.NoError(t, err, "Error al serializar el CandidateDto")

	expectedJSON := `{"id":1,"name":"John Doe","email":"john.doe@example.com","gender":"Male","salary":50000}`
	assert.JSONEq(t, expectedJSON, string(jsonData), "El JSON serializado no coincide con el esperado")
}

func TestCandidateDtoDeserialization(t *testing.T) {
	jsonData := `{"id":1,"name":"John Doe","email":"john.doe@example.com","gender":"Male","salary":50000}`

	var candidate CandidateDto
	err := json.Unmarshal([]byte(jsonData), &candidate)
	assert.NoError(t, err, "Error al deserializar el CandidateDto")

	assert.Equal(t, 1, candidate.ID, "El ID no es correcto")
	assert.Equal(t, "John Doe", candidate.Name, "El nombre no es correcto")
	assert.Equal(t, "john.doe@example.com", candidate.Email, "El email no es correcto")
	assert.Equal(t, "Male", candidate.Gender, "El género no es correcto")
	assert.Equal(t, 50000.00, candidate.Salary, "El salario no es correcto")
}

func TestCandidateDtoSerializationAndDeserialization(t *testing.T) {

	candidate := CandidateDto{
		ID:     2,
		Name:   "Jane Smith",
		Email:  "jane.smith@example.com",
		Gender: "Female",
		Salary: 60000.00,
	}

	jsonData, err := json.Marshal(candidate)
	assert.NoError(t, err, "Error al serializar el CandidateDto")

	var deserializedCandidate CandidateDto
	err = json.Unmarshal(jsonData, &deserializedCandidate)
	assert.NoError(t, err, "Error al deserializar el CandidateDto")

	assert.Equal(t, candidate.ID, deserializedCandidate.ID, "El ID no coincide después de la deserialización")
	assert.Equal(t, candidate.Name, deserializedCandidate.Name, "El nombre no coincide después de la deserialización")
	assert.Equal(t, candidate.Email, deserializedCandidate.Email, "El email no coincide después de la deserialización")
	assert.Equal(t, candidate.Gender, deserializedCandidate.Gender, "El género no coincide después de la deserialización")
	assert.Equal(t, candidate.Salary, deserializedCandidate.Salary, "El salario no coincide después de la deserialización")
}
