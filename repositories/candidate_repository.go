package repositories

import (
	"crud-candidates/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCandidateRepository(db *gorm.DB) *CandidateDao {
	return &CandidateDao{
		DB: db,
	}
}

type CandidateRepository interface {
	CreateCandidate(c *gin.Context, candidate *models.Candidates) error
	GetCandidatesDuplicate(c *gin.Context, candidates *[]models.Candidates, email string) error
	GetCandidates(c *gin.Context, candidates *[]models.Candidates) (*[]models.Candidates, error)
	GetCandidate(c *gin.Context, candidate *models.Candidates) (*models.Candidates, error)
	UpdateCandidate(c *gin.Context, candidate *models.Candidates) error
	DeleteCandidate(c *gin.Context, candidate *models.Candidates) error
}

type CandidateDao struct {
	DB *gorm.DB
}

func (dao *CandidateDao) CreateCandidate(c *gin.Context, candidate *models.Candidates) error {
	if err := dao.DB.Create(candidate).Error; err != nil {
		return err
	}
	return nil
}

func (dao *CandidateDao) GetCandidatesDuplicate(c *gin.Context, candidates *[]models.Candidates, email string) error {
    if err := dao.DB.Where("email = ?", email).Find(candidates).Error; err != nil {
        return err
    }

    return nil
}

func (dao *CandidateDao) GetCandidates(c *gin.Context, candidates *[]models.Candidates) (*[]models.Candidates, error) {
	if err := dao.DB.Find(candidates).Error; err != nil {
		return nil, err
	}
	return candidates, nil
}

func (dao *CandidateDao) GetCandidate(c *gin.Context, candidate *models.Candidates) (*models.Candidates, error) {
	id := c.Param("id")
	if err := dao.DB.First(candidate, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("candidato no encontrado")
		}
		return nil, err
	}
	return candidate, nil
}

func (dao *CandidateDao) UpdateCandidate(c *gin.Context, candidate *models.Candidates) error {
	id := c.Param("id")
	var existingCandidate models.Candidates
	if err := dao.DB.First(&existingCandidate, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("candidato no encontrado")
		}
		return err
	}

	if err := dao.DB.Model(&existingCandidate).Updates(candidate).Error; err != nil {
		return err
	}
	return nil
}

func (dao *CandidateDao) DeleteCandidate(c *gin.Context, candidate *models.Candidates) error {
	id := c.Param("id")
	if err := dao.DB.Delete(candidate, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
