package bootstrap

import (
	db "github.com/pocikode/simple_bank_go/db/sqlc"
)

type Application struct {
	DB *db.Store
}

func App() *Application {
	app := Application{}
	app.DB = NewDatabase()

	return &app
}
