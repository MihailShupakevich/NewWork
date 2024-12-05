package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
	"log"
)

func main() {
	connStr := "user=guest password=guest123 dbname=migrations host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка при проверке подключения: %v", err)
	}
	fmt.Println("Подключение к базе данных успешно установлено.")
}
