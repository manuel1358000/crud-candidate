package routers

import (
    "crud-candidates/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterCandidateRouters(r *gin.Engine, controller *controllers.CandidateController) {
    candidates := r.Group("/candidates")
    {
        candidates.POST("/", controller.CreateCandidate)
        candidates.GET("/", controller.GetCandidates)
        candidates.GET("/:id", controller.GetCandidate)
        candidates.PUT("/:id", controller.UpdateCandidate)
        candidates.DELETE("/:id", controller.DeleteCandidate)
    }
}