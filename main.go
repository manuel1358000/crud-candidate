package main

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "crud-candidates/docs"

    "crud-candidates/clients"
    "crud-candidates/controllers"
    "crud-candidates/middlewares"
    "crud-candidates/repositories"
    "crud-candidates/routers"
)

func main() {
    db := clients.InitDB()

    candidateRepo := repositories.NewCandidateRepository(db)
    candidateController := controllers.NewCandidateController(candidateRepo)
    healthController := controllers.NewHealthController()

    r := gin.Default()
    r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

    r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
        URL: "/swagger.yaml",
    }, swaggerFiles.Handler))

    r.GET("/ping", healthController.Ping)

	r.POST("/login", controllers.LoginController)

    protected := r.Group("/candidates")
    protected.Use(middlewares.JWTMiddleware())
    {
        routers.RegisterCandidateRouters(protected, candidateController)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}
