### Estrutura do diretório Go

Go por padrão utiliza a pasta `~/go` com três subpastas principais:

* **Pkg**: Armazena dependências instaladas para evitar redundância entre projetos.
* **Src**: Local onde o código-fonte é organizado por repositórios.
* **Bin**: Armazena os binários compilados através do comando `go install`.

### Gopath vs Goroot

* **Goroot**: Localização da instalação do binário e bibliotecas padrão do Go.
* **Gopath**: Local do seu workspace, incluindo seu código e pacotes baixados.

### Pacotes vs Modules

* **Pacotes**: Cada diretório individual no projeto é um pacote.
* **Modules**: Coleção de pacotes gerenciada pelo arquivo `go.mod` na raiz.

### Estrutura do go.mod

* **module**: Define o nome e caminho do repositório.
* **go**: Indica a versão do Go utilizada no projeto.
* **require**: Lista dependências e suas versões específicas.
* **replace**: Substitui um módulo remoto por um caminho local ou versão alternativa.
* **exclude**: Impede que versões específicas de pacotes sejam utilizadas.
* **//indirect**: Indica dependência de outra dependência ou módulo não utilizado no código.
* **Go.sum**: Arquivo autogerado com hashes de integridade para commits e pacotes.

### Versões dos Modules

As versões podem ser gerenciadas através de:

* Caminhos internos dentro do próprio módulo.
* Referências diretas via branches.
* Mapeamento via Git tags (ex: 1.0 para v1.0).

### Pseudo-versões

Formatação de pre-release para revisões de commits sem tags de versão semântica:

1. **Prefixo**: Versão base derivada de uma tag anterior ou v0.0.0.
2. **Timestamp**: Data e hora UTC da criação da revisão (aaaammddhhmmss).
3. **Identificador**: Prefixo de 12 caracteres do hash do commit.

### Go Get vs Go Install

* **go get**: Utilizado para gerenciar e baixar dependências no arquivo `go.mod`. O código-fonte é baixado para a pasta `pkg`.
* **go install**: Utilizado para compilar o código e gerar arquivos executáveis. O binário resultante é armazenado na pasta `bin`.

Gostaria que eu demonstrasse como remover uma dependência do seu `go.mod` utilizando o comando `go get`?