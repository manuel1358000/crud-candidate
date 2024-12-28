package routers

import (
    "crud-candidates/controllers"
    "github.com/gin-gonic/gin"
)

func RegisterCandidateRouters(rg *gin.RouterGroup, controller *controllers.CandidateController) {
    {
        rg.POST("/", controller.CreateCandidate)
        rg.GET("/", controller.GetCandidates)
        rg.GET("/:id", controller.GetCandidate)
        rg.PUT("/:id", controller.UpdateCandidate)
        rg.DELETE("/:id", controller.DeleteCandidate)
    }
}