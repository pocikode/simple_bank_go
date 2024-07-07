package main

import (
	"database/sql"
	"log"
	"pocikode/simple-bank/api"
	db "pocikode/simple-bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://default:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8012"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err = server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
