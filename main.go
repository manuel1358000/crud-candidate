package main

import (
	"github.com/gin-gonic/gin"
	"crud-candidates/routers"
	"crud-candidates/clients"
	"crud-candidates/repositories"
	"crud-candidates/controllers"
	"github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    _ "crud-candidates/docs"
)


func main(){
	 db := clients.InitDB()

	 candidateRepo := repositories.NewCandidateRepository(db)
 
	 candidateController := controllers.NewCandidateController(candidateRepo)
 
	 r := gin.Default()
	 // Servir archivo swagger.yaml
	 r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	 // Configurar Swagger UI para leer swagger.yaml
	 r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		 URL: "/swagger.yaml", // Ruta al archivo YAML
	 }, swaggerFiles.Handler))
 
	 routers.RegisterCandidateRouters(r, candidateController)
 
	 r.Run(":8080")
}