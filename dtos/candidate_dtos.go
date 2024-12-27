package dtos


type CandidateDto struct {
	ID        int   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender	  string `json:"gender"`
	Salary	  float64 `json:"salary"`
}

