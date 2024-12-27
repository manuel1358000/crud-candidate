package main

import (
	"github.com/gin-gonic/gin"
	"crud-candidates/routers"
)


func main(){
	r:= gin.Default()

	routers.RegisterCandidateRouters(r)

	r.Run(":8080")
}