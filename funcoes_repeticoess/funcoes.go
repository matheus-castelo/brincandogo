package main

import (
	"errors"
	"fmt"
)

func main() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While
	contador := 0
	for contador <= 10 {
		fmt.Println("MENOR QUE 10")
		contador++
	}

	// Do While
	anExpression := false
	for ok := true; ok; ok = anExpression {
		fmt.Println("OOI")
	}

	// For com Range
	tempArray := []string{"temp1", "temp2", "temp3"} 
	for i, value := range tempArray {
		println(i, value)
	}

	// Ignorando Indice
	for _, value := range tempArray {
		println(value)
	}

	value, codigo := teste()
	fmt.Println(value, codigo)

	valor, err := teste2(false)

	if err != nil {
		fmt.Println("Erro encontrado:", err)
	} else {
		fmt.Println("Sucesso:", valor)
	}

	value, _ = testeRetornoNomeado()
	fmt.Println("Retorno Nomeado:", value)

	// Funcao Anonima atribuida a uma variavel
	funcaoTeste := func(test string, testInt int) {
		fmt.Println(test, testInt)
	}

	// Chamando a função temp e passando a funcaoTeste como argumento
	temp(funcaoTeste)

	outraFuncao := tempFuncao()
	outraFuncao("oi", 4)

}

func teste() (string, int) {
	return "teste", 777
}

func teste2(sucesso bool) (string, error) {
	if !sucesso {
		return "", errors.New("falha catastrófica no sistema")
	}
	return "Dados processados", nil
}

func testeRetornoNomeado() (retornoString string, retornoError error) {
	retornoError = nil
	retornoString = "Isso eh um retorno nomeado"
	return // golang sabe que voce ta retornando essas duas coisas ja
}

// Esta funcao recebe outra funcao como parametro
func temp(funcaoRecebida func(string, int)) {
	funcaoRecebida("otavio", 20)
}

func tempFuncao () func(string,int){
	funcaoTeste := func(valorString string, valorInt int){
		println(valorString, valorInt)
	}
	return funcaoTeste
}