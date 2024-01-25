package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pocikode/simple_bank_go/internal/domain"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			var domainError *domain.Error
			switch {
			case errors.As(err.Err, &domainError):
				ctx.JSON(domainError.Code, gin.H{
					"error": domainError.Message,
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
			}
		}
	}
}
