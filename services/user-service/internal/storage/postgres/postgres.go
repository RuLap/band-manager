package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitDB(connStr string) {
	var err error
	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
}

func GetDB() *pgxpool.Pool {
	return db
}
