package main

import (
	"errors"
	"fmt"
)

func main() {
	if _, err := test(); err != nil {
		fmt.Println("Erro simples:", err.Error())
	}

	if _, err := testPersonalizado(); err != nil {
		fmt.Println("Erro customizado:", err.Error())
	}

	if err := test2(); err != nil {
        fmt.Println(err.Code, err.Message)
		fmt.Printf("Com %%#v: %#v\n", err) // Printando o objeto inteiro - Tipo e estrutura
    }
}

func test() (string, error) {
	return "", errors.New("Erro teste")
}

func testPersonalizado() (string, error) {
	return "", ErrorTest{
		Code:    500,
		Message: "Falha interna no servidor",
	}
}

// Erros personalizados
type ErrorTest struct {
	Code    int
	Message string
}

// Implementação da interface error: func Error() string
func (e ErrorTest) Error() string {
	return fmt.Sprintf("Erro ao processar dados: code=%d, message=%s", e.Code, e.Message)
}


// Posso retornar um PONTEIRO DE ERRORTEST ao invés de ERROR, já que são interfaces
func test2 () *ErrorTest {
	err := &ErrorTest {
		Code: 500,
		Message: "Agora fudeu",
	}
	return err
}