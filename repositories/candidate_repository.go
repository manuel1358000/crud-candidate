package repositories


import(
	"crud-candidates/models"
)

type CandidateRepository interface {
	CreateCandidate(c *gin.Context, candidate *models.Candidate) error
	GetCandidatesDuplicate(c *gin.Context, candidates *models.Candidate) (*models.ListOfCandidates, error)
	GetCandidates(c *gin.Context, candidates *[]models.Candidate) (*[]models.Candidate, error)
	GetCandidate(c *gin.Context, 	candidate *models.Candidate) (*models.Candidate, error)
	UpdateCandidate(c *gin.Context, candidate *models.Candidate) error
	DeleteCandidate(c *gin.Context, candidate *models.Candidate) error
}


type CandidateDao struct {
	DB *gorm.DB
}



func (c *CandidateDao) CreateCandidate(c *gin.Context, candidate *models.Candidate) error {
	c.DB.Create(candidate)
	return nil
}

func (c *CandidateDao) GetCandidatesDuplicate(c *gin.Context, candidates *[]models.Candidate) (*[]models.Candidate, error) {
	c.DB.Find(candidates)
	return candidates, nil
}

func (c *CandidateDao) GetCandidates(c *gin.Context, candidates *[]models.Candidate) (*[]models.Candidate, error) {
	c.DB.Find(candidates)
	return candidates, nil
}	

func (c *CandidateDao) GetCandidate(c *gin.Context, candidate *models.Candidate) (*models.Candidate, error) {
	c.DB.First(candidate, c.Param("id"))
	return candidate, nil
}

func (c *CandidateDao) UpdateCandidate(c *gin.Context, candidate *models.Candidate) error {
	c.DB.Save(candidate)
	return nil
}
func (c *CandidateDao) DeleteCandidate(c *gin.Context, candidate *models.Candidate) error {
	c.DB.Delete(candidate)
	return nil
}