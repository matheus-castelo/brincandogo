package main

import "fmt"

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewPeople(name string, age int) (*People, error) {
	if age < 0 {
		return nil, &ApiErr{
			Name:    "invalid_data",
			Code:    400,
			Message: "Idade não pode ser negativa",
			Causes:  []Causes{{Field: "Age", Message: "valor abaixo de zero"}},
		}
	}
	return &People{Name: name, Age: age}, nil
}

func GetPessoa(p *People) (*People, error) {
	if p == nil {
		return nil, NewNotFoundError("Pessoa não encontrada")
	}
	return p, nil
}

func ChangeName(p *People, newName string) (*People, error) {
	if p == nil {
		return nil, NewNotFoundError("Pessoa inexistente")
	}
	if newName == "" {
		return nil, &ApiErr{Name: "Validation Error", Code: 400, Message: "Nome vazio"}
	}
	p.Name = newName
	return p, nil
}

func DeletePeople(p *People) (string, error) {
	if p == nil {
		return "", NewNotFoundError("Pessoa não encontrada para deletar")
	}
	return fmt.Sprintf("Você deletou %s", p.Name), nil
}