package http

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pocikode/simple_bank_go/internal/domain"
	"github.com/pocikode/simple_bank_go/internal/usecase"
	"net/http"
)

type AccountController struct {
	UseCase *usecase.AccountUseCase
}

func NewAccountController(useCase *usecase.AccountUseCase) *AccountController {
	return &AccountController{UseCase: useCase}
}

func (c *AccountController) Create(ctx *gin.Context) {
	var req domain.CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	account, err := c.UseCase.Create(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (c *AccountController) Get(ctx *gin.Context) {
	var req domain.GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	account, err := c.UseCase.Get(&req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.Error(domain.NewError(http.StatusNotFound, err.Error()))
		} else {
			ctx.Error(err)
		}
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (c *AccountController) List(ctx *gin.Context) {
	var req domain.ListAccountParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.Error(domain.NewError(http.StatusBadRequest, err.Error()))
		return
	}

	accounts, err := c.UseCase.List(&req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
