package main

import (
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"
    "github.com/swaggo/files" 
    _ "example/hello/api_gin/docs" 
)

type Database struct {
	Pessoa []People
}

// @title           People Management API
// @version         1.0
// @description     API para aprendizado de Go e gerenciamento de pessoas.
// @host            localhost:8080
// @BasePath        /
func main() {
	database := &Database{
		Pessoa: []People{
			{Name: "Joaozinho", Age: 20},
		},
	}

	server := &Server{DB: database}
	r := gin.Default()

	// Rotas do CRUD
	r.POST("/people", server.CreateHandler)
	r.GET("/people/:name", server.GetHandler)
	r.PATCH("/people/:name", server.UpdateHandler)
	r.DELETE("/people/:name", server.DeleteHandler)

	// Rota do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}