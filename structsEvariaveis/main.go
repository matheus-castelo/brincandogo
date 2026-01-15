package structEvariaveis

import "fmt"


func main(){
	var teste string = "isso eh uma string"
	outraString := "string"
	fmt.Println(teste)
	fmt.Printf("%T",outraString)

	var temp interface{} // Pode receber qualquer valor.
	temp = "string"
	println(temp)
	temp = 20
	println(temp)

	// Array = Lista tamanho definiido. Slice = Lista tamanho indefinido
	var array1 [4]string = [4]string{"teste", "teste", "teste", "teste"} // Se eu tirar esse 4, ele muda de array para slice
 	println(len(array1))
	println(cap(array1))
	// array1 = append(array1, "otavio") - ERRO. Se fosse um SLICE podia

	usandoUser()
}

// Se a primeira letra da função for maiuscula, ela é publica, se não, privada
func Teste(){
}

// Structs em Golang sao identicas ao que definimos como objetos em outras linguagens.

type User struct {
	name string
	age int 
}

func usandoUser () {
	var user User = User {
		name: "otavio",
		age: 18,
	}

	println("Printando User: ")
	fmt.Println(user)
}