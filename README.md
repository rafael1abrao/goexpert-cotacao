# 💱 Desafio GoExpert: Cotação do Dólar

Este projeto implementa dois serviços em Go — **cliente** e **servidor** — que se comunicam via HTTP. Ele utiliza **contextos com timeout**, **persistência em SQLite** e **manipulação de arquivos** como parte de um desafio prático da formação Goexpert da Fullcycle

---

## 🗂 Estrutura
cotacao/
├── client/
│ └── main.go # Cliente: consome cotação e salva em arquivo
├── server/
│ └── main.go # Servidor: expõe endpoint e grava cotação no banco
├── cotacao.db # Banco SQLite (criado automaticamente)
├── cotacao.txt # Arquivo de saída do cliente
├── go.mod
├── go.sum
└── .gitignore


---

## 🚀 Como executar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/rafael1abrao/goexpert-cotacao.git
cd goexpert/cotacao

go mod tidy

go run ./server

go run ./client

