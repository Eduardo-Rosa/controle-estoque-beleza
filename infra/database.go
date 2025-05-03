package infra

import (
	"database/sql"
	"sync"
	"log"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://user:password@localhost:5432/estoque?sslmode=disable")
		if err != nil {
			log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
		}
	})
	return db
}