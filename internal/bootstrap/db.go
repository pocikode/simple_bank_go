package bootstrap

import (
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/pocikode/simple_bank_go/db/sqlc"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://default:secret@127.0.0.1:5432/simple_bank?sslmode=disable"
)

func NewDatabase() *db.Store {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return db.NewStore(conn)
}
