package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Глобальная переменная для хранения подключения к базе данных
var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/todo-list?sslmode=disable"
)

// Функция инициализации тестового подключения к базе данных
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Создаем новый экземпляр Queries для тестового подключения
	testQueries = New(conn)

	// Запускаем все тесты и сообщаем о результате, затем закрываем подключение к базе данных
	os.Exit(m.Run())
}
