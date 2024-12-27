package dtos


type CandidateDto struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender	  string `json:"gender"`
	Salary	  float64 `json:"salary"`
}

