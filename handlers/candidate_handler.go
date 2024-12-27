package handlers

import (
    "crud-candidates/repositories"
    "github.com/gin-gonic/gin"
    "net/http"
)


type CandidateController struct {   
    repo *repositories.CandidateRepository
}

func NewItemController(repo *repositories.CandidateRepository) *CandidateController {
    return &CandidateController{repo: repo} 
}

func (controller *CandidateController) CreateCandidate(c *gin.Context){


    c.JSON(http.StatusOK, gin.H{ "message": "Create Candidate" })
}

func (controller *CandidateController) GetCandidates(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{ "message": "Get Candidates" })
}

func (controller *CandidateController) GetCandidate(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{ "message": "Get Candidate" })
}

func (controller *CandidateController) UpdateCandidate(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{ "message": "Update Candidate" })
}

func (controller *CandidateController) DeleteCandidate(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{ "message": "Delete Candidate" })
}