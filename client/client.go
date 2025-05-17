// client.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverURL     = "http://localhost:8080/cotacao"
	clientTimeout = 300 * time.Millisecond
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Erro ao chamar servidor: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status inválido do servidor: %s", resp.Status)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	bid := result["bid"]
	if bid == "" {
		log.Fatal("Campo 'bid' vazio na resposta")
	}

	err = salvarEmArquivo(bid)
	if err != nil {
		log.Fatalf("Erro ao salvar no arquivo: %v", err)
	}

	fmt.Printf("Cotação salva: Dólar: %s\n", bid)
}

func salvarEmArquivo(valor string) error {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("Dólar: %s\n", valor))
	return err
}
