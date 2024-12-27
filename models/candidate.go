package models


import(
	"crud-candidates/dtos"
	"github.com/segmentio/ksuid"
)

type Candidates struct {
	ID        uint   `gorm:"primary_key"`
	Name      string
	Email     string
	Gender    string
	Salary    float64
}


type ListOfCandidates []*Candidates


func (candidate *Candidates) MapCandidateDtoFromCandidate() *dtos.CandidateDto {
	return &dtos.CandidateDto{
		ID:        candidate.ID,
		Name:      candidate.Name,
		Email:     candidate.Email,
		Gender:    candidate.Gender,
		Salary:    candidate.Salary,
	}
}


func (listOfCandidates ListOfCandidates) MapCandidateDtoFromCandidates() []dtos.CandidateDto {
	var candidateDtos []*dtos.CandidateDto
	for _, candidate := range listOfCandidates {
		candidateDtos = append(candidateDtos, candidate.MapCandidateDtoFromCandidate())
	}
	return candidateDtos
}


func NewCandidate(candidateDto *dtos.CandidateDto) *Candidates {
	id := ksuid.New().String()
	return &Candidates{
		ID:		id,
		Name:      candidateDto.Name,
		Email:     candidate.Email,
		Gender:    candidate.Gender,
		Salary:   candidate.Salary,
	}
}