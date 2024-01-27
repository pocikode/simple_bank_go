package route

import (
	"github.com/gin-gonic/gin"
	db "github.com/pocikode/simple_bank_go/db/sqlc"
	"github.com/pocikode/simple_bank_go/internal/delivery/http/middleware"
)

func Setup(db db.Store, gin *gin.Engine) {
	gin.Use(middleware.ErrorHandler())

	NewAccountRouter(db, gin)
}
