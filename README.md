# GoLang

Repositório de estudos GoLang.

## Para instalar o go no ubuntu

### Baixa o arquivo com o curl

```bash
curl -OL https://golang.org/dl/go1.20.1.linux-amd64.tar.gz
```

### Executa o comando abaixo para verificar a integridade do arquivo

```bash
sha256sum go1.20.1.linux-amd64.tar.gz
```

### Descompacta o arquivo na /usr/local

```bash
sudo tar -C /usr/local -xvf go1.20.1.linux-amd64.tar.gz
```

### Insere a variável de ambiente no zshrc

```bash
nano ~/.zshrc

# Cria pasta $HOME/goVenzel
# Insere
export GOROOT=/usr/local/go
export GOPATH=$HOME/goVenzel
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

### No Vscode

```json
{
    "go.gopath": "/usr/local/go/bin"
	"[go]": {
        "editor.insertSpaces": true,
        "editor.formatOnSave": true,
        "editor.defaultFormatter": "golang.go"
    },
}
```

### Atualiza o ZSHRC

```bash
source ~/.zshrc
```

### Verifica se o go foi instalado corretamente

```bash
go version
```

## Criar, rodar, build e executar um arquivo main.go

### Criar

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("ok!")
}
```

### Rodar

```bash
go run main.go
```

### Build

```bash
go build main.go
```

### Executar

```bash
./main
```

## VS Code extensões

-   [Go] - Go Team at Google (Possui o intelisense e etc)

## Adicionar um módulo gamer

```bash
go mod init gamer
```
