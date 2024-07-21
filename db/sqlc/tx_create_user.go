package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) (err error) {
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return
		}

		err = arg.AfterCreate(result.User)
		return
	})

	return result, err
}
