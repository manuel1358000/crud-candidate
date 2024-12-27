package routers



func RegisterCandidateRouters(r *gin.Engine) {
	r.POST("/candidate", handler.CreateCandidate)
	r.GET("/candidates", handler.GetCandidates)
	r.GET("/candidate/:id", handler.GetCandidate)
	r.PUT("/candidate/:id", handler.UpdateCandidate)
	r.DELETE("/candidate/:id", handler.DeleteCandidate)
}	