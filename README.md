# Desafio GoExpert: Cotação do Dólar

Este projeto implementa dois serviços em Go — **cliente** e **servidor** — que se comunicam via HTTP. Ele utiliza **contextos com timeout**, **persistência em SQLite** e **manipulação de arquivos** como parte de um desafio prático proposto no curso **GoExpert**.

---

## 🗂 Estrutura

```
cotacao/
├── client/
│   └── main.go         # Cliente: consome cotação e salva em arquivo
├── server/
│   └── main.go         # Servidor: expõe endpoint e grava cotação no banco
├── cotacao.db          # Banco SQLite (criado automaticamente)
├── cotacao.txt         # Arquivo de saída do cliente
├── go.mod
├── go.sum
└── .gitignore
```

---

## 🚀 Como executar o projeto

### 1. Clonar o repositório

```bash
git clone https://github.com/rafael1abrao/goexpert-cotacao.git
cd goexpert-cotacao
```

---

### 2. Instalar dependências

```bash
go mod tidy
```

---

### 3. Rodar o servidor

```bash
go run ./server
```

- Inicia o servidor HTTP na porta `8080`
- Endpoint: `GET /cotacao`
- Consulta a API externa: [https://economia.awesomeapi.com.br/json/last/USD-BRL](https://economia.awesomeapi.com.br/json/last/USD-BRL)
- Persiste no banco SQLite (`cotacao.db`) com timeout de 10ms
- Timeout máximo para chamada da API: **200ms**

---

### 4. Rodar o cliente

Em outro terminal:

```bash
go run ./client
```

- Requisição HTTP ao servidor `http://localhost:8080/cotacao`
- Timeout máximo da requisição: **300ms**
- Extrai o valor `bid` e salva em `cotacao.txt` no formato:
  ```
  Dólar: 5.1234
  ```

---

## 🛢 Consultar o banco SQLite

Você pode visualizar as cotações salvas usando:

```bash
sqlite3 cotacao.db "SELECT * FROM cotacoes;"
```

Ou interativamente:

```bash
sqlite3 cotacao.db
.headers on
.mode column
SELECT * FROM cotacoes;
.quit
```

---

## ⚠️ Timeouts utilizados

| Componente        | Operação            | Timeout |
|------------------|---------------------|---------|
| `server/main.go` | Consulta à API      | 200ms   |
| `server/main.go` | Gravação no banco   | 10ms    |
| `client/main.go` | Requisição HTTP     | 300ms   |

Todos os contextos utilizam `context.Context` com `timeout`, e erros são registrados em log quando o tempo de execução é excedido.

---

## ✅ Requisitos do Desafio

- [x] `server.go` escutando na porta 8080
- [x] Endpoint `/cotacao` disponível
- [x] Chamada à API externa com timeout de 200ms
- [x] Persistência em banco com timeout de 10ms
- [x] Cliente com timeout de 300ms
- [x] Arquivo `cotacao.txt` contendo apenas o valor da cotação
- [x] Uso consistente do package `context`

---

## 📋 Tecnologias utilizadas

- [Go](https://golang.org/)
- [SQLite](https://sqlite.org/)
- [AwesomeAPI - Dólar](https://economia.awesomeapi.com.br/)
- `context`, `http`, `os`, `encoding/json`, `database/sql`, `log`, etc.

---

## 👨‍💻 Autor

Desenvolvido por [Rafael Abrão](https://github.com/rafael1abrao) como parte do curso **GoExpert** da **FullCycle**.

---
