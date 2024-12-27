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

	 healthController := controllers.NewHealthController()
 
	 r := gin.Default()
	 r.StaticFile("/swagger.yaml", "./docs/swagger.yaml")

	 r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(&ginSwagger.Config{
		 URL: "/swagger.yaml", 
	 }, swaggerFiles.Handler))
 
	 r.GET("/ping", healthController.Ping)

	 routers.RegisterCandidateRouters(r, candidateController)
 
	 r.Run(":8080")
}