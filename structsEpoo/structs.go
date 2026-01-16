package main

import (
	"errors"
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}
// Não temos construtores antivos. Criamos funçõesq ue retornam a struct, geralmente prefixadas com New.
func NewPessoa(nome string, idade int) (*Pessoa, error) {
	if idade < 0 {
		return nil, errors.New("idade nao pode ser negativa")
	}
	
	return &Pessoa{Nome: nome, Idade: idade}, nil
}

func main() {
	pessoa, err := NewPessoa("Joaozinho", 20)

	if err != nil {
		fmt.Println("Erro ao criar pessoa:", err)
		return 
	}

	fmt.Printf("Pessoa criada com sucesso: %+v\n", *pessoa)


	println(pessoa.Idade)
	pessoa.FazerAniversario()
	println(pessoa.Idade)
	pessoa.Saudacao()


	dev1 := Desenvolvedor{
		Pessoa: *pessoa,
		Linguagem: "Go",
	}
	println(dev1.Nome)
	println(dev1.Linguagem)
	dev1.FazerAniversario()
	println(dev1.Idade)
	dev1.Saudacao()

	// Saída para mostrar exemplo de interfaces
	IniciarJornada(dev1)
	// IniciarJornada(pessoa) Não posso iniciar jornada com a pessoa. Porque pessoa não tem habilidade de trabalhar. Mesmo que Desenvolvedor seja uma pessoa, o contrário não é verdade.


	
	
}

/* 
Os métodos em Go não ficam em classes. Fazemos de duas maneiras

- (T) Value Receiver = Recebe uma cópia da Struct (T) (OBS: por consequencia, nao altera nada, pois não temos o endereço da memória da struct original)

- (*T) Pointer Receiver = Recebe a referência (*T)
*/

// Pointer Receiver
func (p *Pessoa) FazerAniversario(){
	p.Idade++
}

// Value Receiver

func (p Pessoa) Saudacao() string {
	return "Olá, meu nome é " + p.Nome
}


// Golang não possui herança. Nós passamos outra struct para uma
type Desenvolvedor struct {
	Pessoa
	Linguagem string
}


// Interfaces. Se minha vó tivesse barba seria meu vô.  A implementação de Interfaces é implícita. Se a struct tem os métodos que ai nterface pede, ela'e aquela interface.
type Trabalhador interface {
	Trabalhar()
}

func (d Desenvolvedor) Trabalhar(){
	println(d.Nome, "esta codando em", d.Linguagem)
}

func IniciarJornada(t Trabalhador){
	t.Trabalhar()
}