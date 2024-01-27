package repository

import (
	"context"
	db "github.com/pocikode/simple_bank_go/db/sqlc"
)

type AccountRepository struct {
	db db.Store
}

func NewAccountRepository(db db.Store) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Create(args *db.CreateAccountParams) (*db.Account, error) {
	result, err := ar.db.CreateAccount(context.Background(), *args)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ar *AccountRepository) FindByID(id int64) (*db.Account, error) {
	result, err := ar.db.GetAccount(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ar *AccountRepository) List(args *db.ListAccountParams) (*[]db.Account, error) {
	result, err := ar.db.ListAccount(context.Background(), *args)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
