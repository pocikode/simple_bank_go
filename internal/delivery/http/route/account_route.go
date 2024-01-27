package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/pocikode/simple_bank_go/db/sqlc"
	"github.com/pocikode/simple_bank_go/internal/delivery/http"
	"github.com/pocikode/simple_bank_go/internal/repository"
	"github.com/pocikode/simple_bank_go/internal/usecase"
)

func NewAccountRouter(db db.Store, gin *gin.Engine) {
	ar := repository.NewAccountRepository(db)
	auc := usecase.NewAccountUseCase(ar)
	ac := http.NewAccountController(auc)

	gin.POST("/accounts", ac.Create)
	gin.GET("/accounts/:id", ac.Get)
	gin.GET("/accounts", ac.List)
}
