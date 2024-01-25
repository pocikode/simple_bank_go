package bootstrap

import (
	db "github.com/pocikode/simple_bank_go/db/sqlc"
)

type Application struct {
	DB  *db.Store
	Env *Env
}

func App() *Application {
	app := Application{}
	app.Env = NewEnv()
	app.DB = NewDatabase(app.Env)

	return &app
}
