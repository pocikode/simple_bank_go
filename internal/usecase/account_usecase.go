package usecase

import (
	db "github.com/pocikode/simple_bank_go/db/sqlc"
	"github.com/pocikode/simple_bank_go/internal/domain"
	"github.com/pocikode/simple_bank_go/internal/repository"
)

type AccountUseCase struct {
	repo *repository.AccountRepository
}

func NewAccountUseCase(repo *repository.AccountRepository) *AccountUseCase {
	return &AccountUseCase{repo: repo}
}

func (uc *AccountUseCase) Create(request *domain.CreateAccountRequest) (*db.Account, error) {
	return uc.repo.Create(&db.CreateAccountParams{
		Owner:    request.Owner,
		Currency: request.Currency,
		Balance:  0,
	})
}

func (uc *AccountUseCase) Get(request *domain.GetAccountRequest) (*db.Account, error) {
	return uc.repo.FindByID(request.ID)
}

func (uc *AccountUseCase) List(request *domain.ListAccountParams) (*[]db.Account, error) {
	return uc.repo.List(&db.ListAccountParams{
		Limit:  request.Size,
		Offset: (request.Page - 1) * request.Size,
	})
}
