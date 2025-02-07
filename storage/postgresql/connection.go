package postgresql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New() (Storage, error) {
	connStr := "user=postgres dbname=blog sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)

	return Storage{DB: db}, err

}
