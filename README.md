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

# Insere
export PATH=$PATH:/usr/local/go/bin
```

### Atualiza o ZSHRC

```bash
source ~/.zshrc
```

### Verifica se o go foi instalado corretamente

```bash
go version
```
