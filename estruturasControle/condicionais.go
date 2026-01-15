package main

func main() {
    if retorno := teste(); retorno == "teste" {
        println("Isso eh um teste")
    }
	// printLn(retonro) - ELA NÃO DA PARA SER ACESSADA
	// if retorno, err := teste(); err != nil {
	//}

	switches()
}

func teste() string {
    return "teste"
}


func switches () { 
	temp := "temp"
	switch temp {
	case "teste", "teste2", "terceiroValor":
			println("Faz isso ...")
		fallthrough // Os switches de outras linguagens checam todos os cases, mesmo que caia no primeiro, ao menos que voce tenha break. Go é o controario, ele sai a partir do momento que dá Verdadeiro, mas pode continuar para o proximo case se voce colocar fallthrough
	case temp:
			println("Faz aquilo...")
		
	case "pararCodigo":
		break

	case "4":

	default : 
	println("Voce nao caiu em nenhum dos 4 cases")
		}
	
}