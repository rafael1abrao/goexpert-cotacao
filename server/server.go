// server.go
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL       = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbName       = "cotacao.db"
	apiTimeout   = 200 * time.Millisecond
	dbTimeout    = 10 * time.Millisecond
	httpAddr     = ":8080"
	endpointPath = "/cotacao"
)

type CotacaoResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db := initDB()
	defer db.Close()

	http.HandleFunc(endpointPath, func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), apiTimeout)
		defer cancel()

		cotacao, err := fetchCotacao(ctx)
		if err != nil {
			http.Error(w, "Erro ao obter cotação", http.StatusRequestTimeout)
			log.Printf("Erro fetch: %v", err)
			return
		}

		ctxDB, cancelDB := context.WithTimeout(r.Context(), dbTimeout)
		defer cancelDB()

		err = salvarCotacao(ctxDB, db, cotacao)
		if err != nil {
			log.Printf("Erro ao salvar no banco: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"bid": cotacao})
	})

	log.Printf("Servidor escutando em %s", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("Erro ao abrir banco:", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		criado_em DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	return db
}

func fetchCotacao(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var resultado CotacaoResponse
	if err := json.NewDecoder(resp.Body).Decode(&resultado); err != nil {
		return "", err
	}

	return resultado.USDBRL.Bid, nil
}

func salvarCotacao(ctx context.Context, db *sql.DB, bid string) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO cotacoes (bid) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, bid)
	return err
}
