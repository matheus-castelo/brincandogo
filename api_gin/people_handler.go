package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	DB *Database
}

func sendError(c *gin.Context, err error) {
	if apiErr, ok := err.(*ApiErr); ok {
		c.JSON(apiErr.Code, apiErr)
		return
	}
	c.JSON(500, gin.H{"message": "internal server error", "error": err.Error()})
}

// CreateHandler godoc
// @Summary      Cria uma nova pessoa
// @Description  Recebe um JSON, valida os dados e armazena no banco em memória
// @Tags         people
// @Accept       json
// @Produce      json
// @Param        pessoa  body      People  true  "Dados da Pessoa"
// @Success      201     {object}  People
// @Failure      400     {object}  ApiErr
// @Router       /people [post]
func (s *Server) CreateHandler(c *gin.Context) {
	var body People
	if err := c.ShouldBindJSON(&body); err != nil {
		sendError(c, NewBadRequestError("Dados invalidos", nil))
		return
	}

	p, err := NewPeople(body.Name, body.Age)
	if err != nil {
		sendError(c, err)
		return
	}

	s.DB.Pessoa = append(s.DB.Pessoa, *p)
	c.JSON(201, p)
}

// GetHandler godoc
// @Summary      Busca uma pessoa
// @Description  Retorna os detalhes de uma pessoa através do nome fornecido na URL
// @Tags         people
// @Produce      json
// @Param        name  path      string  true  "Nome da Pessoa"
// @Success      200   {object}  People
// @Failure      404   {object}  ApiErr
// @Router       /people/{name} [get]
func (s *Server) GetHandler(c *gin.Context) {
	name := c.Param("name")
	var foundPerson *People

	for i := range s.DB.Pessoa {
		if s.DB.Pessoa[i].Name == name {
			foundPerson = &s.DB.Pessoa[i]
			break
		}
	}

	pessoa, err := GetPessoa(foundPerson)
	if err != nil {
		sendError(c, err)
		return
	}

	c.JSON(200, pessoa)
}

// UpdateHandler godoc
// @Summary      Atualiza o nome de uma pessoa
// @Description  Altera o nome de um registro existente no banco em memória
// @Tags         people
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "Nome Atual"
// @Param        body  body      object  true  "Novo Nome"
// @Success      200   {object}  People
// @Failure      400   {object}  ApiErr
// @Failure      404   {object}  ApiErr
// @Router       /people/{name} [patch]
func (s *Server) UpdateHandler(c *gin.Context) {
	name := c.Param("name")
	var foundPerson *People

	for i := range s.DB.Pessoa {
		if s.DB.Pessoa[i].Name == name {
			foundPerson = &s.DB.Pessoa[i]
			break
		}
	}

	var body struct {
		NewName string `json:"new_name"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		sendError(c, NewBadRequestError("JSON inválido", nil))
		return
	}

	updatedPerson, err := ChangeName(foundPerson, body.NewName)
	if err != nil {
		sendError(c, err)
		return
	}

	c.JSON(200, updatedPerson)
}

// DeleteHandler godoc
// @Summary      Remove uma pessoa
// @Description  Deleta o registro do slice em memória baseado no nome
// @Tags         people
// @Param        name  path      string  true  "Nome da Pessoa"
// @Success      200   {object}  map[string]string
// @Failure      404   {object}  ApiErr
// @Router       /people/{name} [delete]
func (s *Server) DeleteHandler(c *gin.Context) {
	name := c.Param("name")
	indexFound := -1

	for i := range s.DB.Pessoa {
		if s.DB.Pessoa[i].Name == name {
			indexFound = i
			break
		}
	}

	if indexFound == -1 {
		_, err := DeletePeople(nil)
		sendError(c, err)
		return
	}

	msg, err := DeletePeople(&s.DB.Pessoa[indexFound])
	if err != nil {
		sendError(c, err)
		return
	}

	s.DB.Pessoa = append(s.DB.Pessoa[:indexFound], s.DB.Pessoa[indexFound+1:]...)
	c.JSON(200, gin.H{"message": msg})
}