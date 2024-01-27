package bootstrap

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	db "github.com/pocikode/simple_bank_go/db/sqlc"
	"log"
)

func NewDatabase(env *Env) db.Store {
	dbDriver := "postgres"
	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser,
		env.DBPass,
		env.DBHost,
		env.DBPort,
		env.DBName,
	)

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return db.NewStore(conn)
}
