package controllers

import (
    "crud-candidates/models"
    "crud-candidates/repositories"
    "crud-candidates/dtos"
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

type CandidateController struct {
    repo repositories.CandidateRepository
}

func NewCandidateController(repo repositories.CandidateRepository) *CandidateController {
    return &CandidateController{repo: repo}
}

func (controller *CandidateController) CreateCandidate(c *gin.Context) {
    var candidateDTO dtos.CandidateDto
    if err := c.ShouldBindJSON(&candidateDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var existingCandidate []models.Candidates
    err := controller.repo.GetCandidatesDuplicate(c, &existingCandidate, candidateDTO.Email)
    if err != nil {
        fmt.Println("Error al verificar duplicados:", err)
        c.Error(fmt.Errorf("Error al verificar duplicados: %v", err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar duplicados"})
        return
    }

    if len(existingCandidate) > 0 {
        c.JSON(http.StatusConflict, gin.H{"error": "Ya existe un candidato con el mismo correo o teléfono"})
        return
    }

    candidate := models.Candidates{
        Name:     candidateDTO.Name,
        Email:    candidateDTO.Email,
        Salary:   candidateDTO.Salary,
        Gender:   candidateDTO.Gender,
    }

    if err := controller.repo.CreateCandidate(c, &candidate); err != nil {
        fmt.Println("Error al crear candidato:", err)
        c.Error(fmt.Errorf("Error al verificar duplicados: %v", err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el candidato"})
        return
    }

    c.JSON(http.StatusCreated, candidate)
}

func (controller *CandidateController) GetCandidates(c *gin.Context) {
    var candidates []models.Candidates
    result, err := controller.repo.GetCandidates(c, &candidates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los candidatos"})
        return
    }

    c.JSON(http.StatusOK, result)
}

func (controller *CandidateController) GetCandidate(c *gin.Context) {
    var candidate models.Candidates
    result, err := controller.repo.GetCandidate(c, &candidate)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Candidato no encontrado"})
        return
    }

    c.JSON(http.StatusOK, result)
}

func (controller *CandidateController) UpdateCandidate(c *gin.Context) {
    var candidateDTO dtos.CandidateDto
    if err := c.ShouldBindJSON(&candidateDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    candidate := models.Candidates{
        Name:     candidateDTO.Name,
        Email:    candidateDTO.Email,
        Salary:   candidateDTO.Salary,
        Gender:   candidateDTO.Gender,
    }

    if err := controller.repo.UpdateCandidate(c, &candidate); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el candidato"})
        return
    }

    c.JSON(http.StatusOK, candidate)
}

func (controller *CandidateController) DeleteCandidate(c *gin.Context) {
    var candidate models.Candidates
    if err := controller.repo.DeleteCandidate(c, &candidate); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el candidato"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Candidato eliminado"})
}
