package models


import(
	"crud-candidates/dtos"
)

type Candidates struct {
	ID        int     `gorm:"primaryKey;autoIncrement"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Email     string  `gorm:"type:varchar(255);unique;not null"`
	Gender    string  `gorm:"type:varchar(50)"`
	Salary    float64 `gorm:"type:decimal(10,2);not null"`
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


func (listOfCandidates ListOfCandidates) MapCandidateDtoFromCandidates() []*dtos.CandidateDto {
	var candidateDtos []*dtos.CandidateDto
	for _, candidate := range listOfCandidates {
		candidateDtos = append(candidateDtos, candidate.MapCandidateDtoFromCandidate())
	}
	return candidateDtos
}


func NewCandidate(candidateDto *dtos.CandidateDto) *Candidates {
	return &Candidates{
		Name:      candidateDto.Name,
		Email:     candidateDto.Email,
		Gender:    candidateDto.Gender,
		Salary:   candidateDto.Salary,
	}
}